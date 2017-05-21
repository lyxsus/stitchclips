package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// App is the server App structure. It contains the router and the configuration
type App struct {
	Router *mux.Router
	Config Config
}

var a = App{}

func initialise() {
	a.Config = LoadConfig()
	a.Router = Router()
}

func main() {
	initialise()
	if _, err := os.Stat(a.Config.Path); os.IsNotExist(err) {
		err := os.Mkdir(a.Config.Path, 0777)
		if err != nil {
			log.Printf("Couldn't create %s: %s\n", a.Config.Path, err)
			os.Exit(1)
		}
	}

	log.Println("Starting server on port 8000")
	log.Fatal(http.ListenAndServe(":"+a.Config.Port, Router()))
}
