package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"gopkg.in/mgo.v2"
)

// App is the server App structure. It contains the router and the configuration
type App struct {
	Router *mux.Router
	Config Config
	Dm     DownloadingManager
	Db     *mgo.Database
}

var a = App{}

func init() {
	LoadConfig()
	a.Router = Router()
	a.Dm = CreateDownloadingManager()
	a.Db = CreateSession(a.Config.DBURL, a.Config.DBName)
}

func main() {
	if _, err := os.Stat(a.Config.Path); os.IsNotExist(err) {
		err := os.Mkdir(a.Config.Path, 0755)
		if err != nil {
			log.Printf("Couldn't create %s: %s\n", a.Config.Path, err)
			os.Exit(1)
		}
	}
	log.Println("Starting server on port", a.Config.Port)

	go a.Dm.run()

	handler := cors.Default().Handler(Router())
	log.Fatal(http.ListenAndServe(":"+a.Config.Port, handler))
}
