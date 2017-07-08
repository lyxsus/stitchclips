package main

import (
	"log"
	"sync"
)

type DownloadingManager struct {
	Outs        map[string]chan string
	Errors      map[string]chan error
	In          chan string
	outMutex    sync.Mutex
	errorsMutex sync.Mutex
	setMutex    sync.Mutex
	Sets        map[string]Clips
}

func CreateDownloadingManager() DownloadingManager {
	log.Println("Initializing DownloadingManager")
	dm := DownloadingManager{
		Outs:   make(map[string]chan string),
		Errors: make(map[string]chan error),
		In:     make(chan string),
		Sets:   make(map[string]Clips),
	}
	return dm
}

func (this DownloadingManager) addClips(clips Clips) (chan string, chan error) {
	// TODO: Check if the clips exists in the database before adding to the list; if it does exist, write on Out channel
	this.outMutex.Lock()
	this.Outs[clips.ID] = make(chan string)
	this.outMutex.Unlock()

	this.errorsMutex.Lock()
	this.Errors[clips.ID] = make(chan error)
	this.errorsMutex.Unlock()

	this.setMutex.Lock()
	if _, ok := this.Sets[clips.ID]; ok != true {
		this.Sets[clips.ID] = clips
		this.In <- clips.ID
	}
	this.setMutex.Unlock()
	return this.Outs[clips.ID], this.Errors[clips.ID]
}

func (this DownloadingManager) run() {
	log.Println("Starting Downloading Manager")
	for true {
		clipsID := <-this.In
		clips := this.Sets[clipsID]
		log.Println("[DownloadingManager] Getting and downloading ", clips.Slugs())
		for _, clip := range clips.Clips {
			err := clip.Get()
			if err != nil {
				this.Errors[clipsID] <- err
			} else {
				err = clip.Download()
				log.Println("[DownloadingManager] Finished: ", clip.Slug)
				if err != nil {
					this.Errors[clipsID] <- err
				} else {
					// TODO: add the clip to the database
					this.Outs[clipsID] <- clip.Slug
					this.Errors[clipsID] <- nil
				}
			}
		}
		close(this.Outs[clipsID])

		close(this.Errors[clipsID])

		this.setMutex.Lock()
		delete(this.Sets, clipsID)
		this.setMutex.Unlock()

		this.outMutex.Lock()
		delete(this.Outs, clipsID)
		this.outMutex.Unlock()

		this.errorsMutex.Lock()
		delete(this.Errors, clipsID)
		this.errorsMutex.Unlock()
		log.Println("[DownladingManager] Done with: ", clips.Slugs())
	}
}
