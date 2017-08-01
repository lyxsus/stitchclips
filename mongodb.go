package main

import (
	"gopkg.in/mgo.v2"
	"log"
)

// CreateSession creates a mgo session
func CreateSession(url string, dbName string) *mgo.Database {
	session, err := mgo.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
	db := session.DB(dbName)
	return db
}
