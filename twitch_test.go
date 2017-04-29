package main

import (
	"os"
	"testing"
)

var testClip = Clip{
	Slug: "DarlingBoringWitchKreygasm",
	Thumbnails: Thumbnails{
		Medium: "https://clips-media-assets.twitch.tv/25144387424-offset-11439.988-42-preview-480x272.jpg",
	},
}

func TestDownloadClip(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	testClip.DownloadClip("clips")
	if _, err := os.Stat("clips/" + testClip.Slug + ".mp4"); os.IsNotExist(err) {
		t.Error("File not downloaded")
	}
}

func TestMain(m *testing.M) {
	os.Remove("clips")
	os.Exit(m.Run())
}
