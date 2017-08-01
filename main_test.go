package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	go a.Dm.run()

	err := a.Db.DropDatabase()
	if err != nil {
		panic(err)
	}

	if a.Config.ClientID == "" {
		a.Config.ClientID = os.Getenv("clientId")
	}

	os.RemoveAll(a.Config.Path)
	os.Mkdir(a.Config.Path, 0777)
	os.Exit(m.Run())
}
