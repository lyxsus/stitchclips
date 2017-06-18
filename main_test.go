package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	initialise()

	a.Db.FlushDB()

	if a.Config.ClientID == "" {
		a.Config.ClientID = os.Getenv("clientId")
	}

	os.RemoveAll(a.Config.Path)
	os.Mkdir(a.Config.Path, 0777)
	os.Exit(m.Run())
}
