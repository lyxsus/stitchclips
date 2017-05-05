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

func TestGetTop(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	clips := Clips{}
	clips.GetTop("itmejp", "2", "all")
	if len(clips.Clips) != 2 {
		t.Error("GetTopClips: did not get clips")
	}
	if clips.Clips[0].Broadcaster.Name != "itmejp" {
		t.Error("GetTopClips: got the wrong clips")
	}
}

func TestDownload(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	testClip.Download()
	if _, err := os.Stat("clips_test/" + testClip.Slug + ".mp4"); os.IsNotExist(err) {
		t.Error("DownloadFile: File not downloaded")
	}
}

func TestMain(m *testing.M) {
	if config.ClientID == "" {
		config.ClientID = os.Getenv("clientId")
	}
	os.Remove("clips_test")
	os.Exit(m.Run())
}
