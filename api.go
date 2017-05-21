package main

import (
	"encoding/json"
	"net/http"
	"os"

	"log"

	"io/ioutil"

	"github.com/gorilla/mux"
)

// StitchedVideo is the compilation of clips
type StitchedVideo struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

// Router returns the router containing all the routes for the API
func Router() *mux.Router {
	r := mux.NewRouter()
	r.Host(a.Config.Host)
	r.HandleFunc("/clips/{channel}/{period}/{limit}", HandleGetClips)
	r.HandleFunc("/stitch", HandleStitch).Methods("POST")
	r.HandleFunc("/video/{id}", HandleVideo).Methods("GET")

	return r
}

// HandleVideo return the .MP4 stitched video
func HandleVideo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	http.ServeFile(w, r, a.Config.Path+"/"+vars["id"])
}

// HandleGetClips returns clips depending on parameters
func HandleGetClips(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	clips := Clips{}
	clips.GetTop(vars["channel"], vars["limit"], vars["period"])
	json, err := json.Marshal(clips)
	if err != nil {
		log.Printf("Error on HandleGetClips: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(json)
}

// HandleStitch stiches clips passed as parameters and returns video URL
func HandleStitch(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	stitchingFile := a.Config.Path + "/" + GetUUID()
	log.Printf("Creating stitching file: %s.\n", stitchingFile)
	_, err := os.Create(stitchingFile)
	if err != nil {
		log.Printf("Error on HandleStitch: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	outputFile := GetUUID()
	outputPath := a.Config.Path + "/" + outputFile
	clips := Clips{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error on reading HandleStitch request body: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = json.Unmarshal(body, &clips)
	if err != nil {
		log.Printf("Error on unserializing json: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	done := make(chan bool, len(clips.Clips))
	errors := make(chan error, len(clips.Clips))
	for _, clip := range clips.Clips {
		err := clip.Get()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = clip.Download()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = clip.PrepareAsync(stitchingFile, done, errors)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	for i := 0; i < len(clips.Clips); i++ {
		<-done
		err = <-errors
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	err = clips.Stitch(outputPath, stitchingFile)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	video := StitchedVideo{
		ID:  outputFile,
		URL: a.Config.Host + ":" + a.Config.Port + "/video/" + outputFile + ".mp4",
	}
	json, err := json.Marshal(video)
	if err != nil {
		log.Printf("Error on serializing json: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(json)
}
