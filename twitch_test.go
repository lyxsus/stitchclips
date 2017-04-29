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

func TestGetTopClips(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	clips := Clips{}
	clips.GetTopClips("itmejp", "2", "all")
	if len(clips.Clips) != 2 {
		t.Error("GetTopClips: did not get clips")
	}
	if clips.Clips[0].Broadcaster.Name != "itmejp" {
		t.Error("GetTopClips: got the wrong clips")
	}
}

func TestDownloadClip(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	testClip.DownloadClip("clips")
	if _, err := os.Stat("clips/" + testClip.Slug + ".mp4"); os.IsNotExist(err) {
		t.Error("DownloadFile: File not downloaded")
	}
}

func TestMain(m *testing.M) {
	os.Remove("clips")
	os.Exit(m.Run())
}
