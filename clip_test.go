package main

import (
	"testing"
	"os"
)

func TestDownload(t *testing.T) {
	t.Log("Testing Download")
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	testClip.Download()
	if _, err := os.Stat(a.Config.Path + "/" + testClip.Slug + ".mp4"); os.IsNotExist(err) {
		t.Error("DownloadFile: File not downloaded")
	}
}
