package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// App is the server App structure. It contains the router and the configuration
type App struct {
	Router *mux.Router
	Config Config
	Db     *redis.Client
}

var a = App{}

func initialise() {
	a.Config = LoadConfig()
	a.Router = Router()
	a.Db = CreateClient()
}

func main() {
	initialise()
	if _, err := os.Stat(a.Config.Path); os.IsNotExist(err) {
		err := os.Mkdir(a.Config.Path, 0755)
		if err != nil {
			log.Printf("Couldn't create %s: %s\n", a.Config.Path, err)
			os.Exit(1)
		}
	}
	log.Println("Starting server on port 8000")

	handler := cors.Default().Handler(Router())
	log.Fatal(http.ListenAndServe(":"+a.Config.Port, handler))
}
