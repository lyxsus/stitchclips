package main

import (
	"log"
	"os"
	"testing"
)

func TestStitch(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	t.Log("Testing Stitch")
	clips := Clips{}
	outputFile := config.Path + "/stitched"
	stitchingFile := config.Path + "/stitching"
	_, err := os.Create(stitchingFile)
	if err != nil {
		log.Println("Error creating stitchingFile: ", err)
	}
	clips.GetTop("itmejp", "2", "all")
	for _, clip := range clips.Clips {
		clip.Download()
		clip.ToMPG(stitchingFile)
	}
	clips.Stitch(outputFile, stitchingFile)
	if _, err := os.Stat(config.Path + "/stitched.mp4"); err != nil {
		t.Error("Did not stitch")
	}
}
