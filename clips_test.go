package main

import (
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
	t.Log("Testing GetTop")
	clips := Clips{}
	clips.GetTop("itmejp", "2", "all")
	if len(clips.Clips) != 2 {
		t.Error("GetTopClips: did not get clips")
	}
	if clips.Clips[0].Broadcaster.Name != "itmejp" {
		t.Error("GetTopClips: got the wrong clips")
	}
}

