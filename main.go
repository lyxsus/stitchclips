package main

import (
	"log"
	"net/http"
)

var config = LoadConfig()

func main() {
	log.Println("Starting server on port 8000")
	log.Fatal(http.ListenAndServe(":"+config.Port, Router()))
}
