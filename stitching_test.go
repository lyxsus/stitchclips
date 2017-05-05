package main

import (
	"os"
	"testing"
)

func TestStitch(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	clips := Clips{}
	config.Output = "stitched"
	clips.GetTop("itmejp", "2", "all")
	for _, clip := range clips.Clips {
		clip.Download()
		clip.ToMPG()
	}
	clips.Stitch()
	if _, err := os.Stat("stitched.mp4"); err != nil {
		t.Error("Did not sitch")
	}
}
