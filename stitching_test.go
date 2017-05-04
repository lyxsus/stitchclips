package main

import (
	"os"
	"testing"
)

func TestStitchClips(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	clips := Clips{}
	clips.GetTopClips("itmejp", "2", "all")
	for _, clip := range clips.Clips {
		clip.DownloadClip("clips_test")
	}
	clips.StitchClips("clips_test")
	if _, err := os.Stat("stitched.mp4"); err != nil {
		t.Error("Did not sitch")
	}
}
