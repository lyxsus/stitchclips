package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var config = LoadConfig()

func router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/clips/{channel}/{period}/{limit}", HandleGetClips)
	r.HandleFunc("/stitch", HandleStitch).Methods("POST")

	return r
}

func main() {
	log.Println("Starting server on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router()))

	// clips.GetTop(config.Channel, config.Limit, config.Period)

	// done := make(chan bool, len(clips.Clips))
	// for _, clip := range clips.Clips {
	// 	clip.Download()
	// 	clip.ToMPGAsync(done)
	// }
	// for i := 0; i < len(clips.Clips); i++ {
	// 	<-done
	// }
	// clips.Stitch()
	// for _, clip := range clips.Clips {
	// 	clip.Cleanup()
	// }
}
