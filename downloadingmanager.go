package main

import (
	"log"
	"sync"
)

type DownloadingManager struct {
	In         chan string
	ClipsMutex sync.Mutex
	Clips      map[string]DmClip
}

type DmClip struct {
	Clip        Clip
	outMutex    sync.Mutex
	errorsMutex sync.Mutex
	Outs        map[string]chan string
	Errs        map[string]chan error
}

func (dmClip DmClip) write(message string) {
	for _, channel := range dmClip.Outs {
		channel <- message
	}
}

func (dmClip DmClip) writeError(err error) {
	for _, channel := range dmClip.Errs {
		channel <- err
	}
}

func CreateDownloadingManager() DownloadingManager {
	log.Println("Initializing DownloadingManager")
	dm := DownloadingManager{
		In:    make(chan string, 8192),
		Clips: make(map[string]DmClip),
	}
	return dm
}

func (this DownloadingManager) addClips(clips Clips) (chan string, chan error) {
	// TODO: Check if the clips exists in the database before adding to the list; if it does exist, write on Out channel

	outChan := make(chan string)
	errChan := make(chan error)

	for _, clip := range clips.Clips {
		this.ClipsMutex.Lock()
		if _, ok := this.Clips[clip.Slug]; ok != true {
			this.Clips[clip.Slug] = DmClip{
				Clip: clip,
				Outs: make(map[string]chan string),
				Errs: make(map[string]chan error),
			}
		}
		this.Clips[clip.Slug].Outs[clips.ID] = outChan
		this.Clips[clip.Slug].Errs[clips.ID] = errChan
		this.ClipsMutex.Unlock()

		this.In <- clip.Slug

	}
	return outChan, errChan
}

func (this DownloadingManager) checkChannelAmount(id string) int {
	amount := 0
	for _, dmClip := range this.Clips  {
		if _, ok := dmClip.Outs[id]; ok == true {
			amount++
		}
	}
	return amount
}

func (this DownloadingManager) run() {
	log.Println("Starting Downloading Manager")
	for true {
		id := <-this.In
		dmClip := this.Clips[id]
		log.Println("[DownloadingManager] Getting and downloading ", dmClip.Clip.Slug)
		err := dmClip.Clip.Get()
		if err != nil {
			dmClip.writeError(err)
		} else {
			err = dmClip.Clip.Download()
			log.Println("[DownloadingManager] Finished: ", dmClip.Clip.Slug)
			if err != nil {
				dmClip.writeError(err)
			} else {
				// TODO: add the clip to the database
				dmClip.write(dmClip.Clip.Slug)
				dmClip.writeError(nil)
			}
		}

		if this.checkChannelAmount(id) == 1 {
			close(dmClip.Outs[id])
			close(dmClip.Errs[id])
		}

		this.ClipsMutex.Lock()
		delete(this.Clips, dmClip.Clip.Slug)
		this.ClipsMutex.Unlock()

		log.Println("[DownladingManager] Done with: ", dmClip.Clip.Slug)
	}
}
