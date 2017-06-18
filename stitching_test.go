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
	log.Println("Starting TestStitch test 1")
	clips := Clips{}
	outputFile := a.Config.Path + "/stitched"
	stitchingFile := a.Config.Path + "/stitching"
	_, err := os.Create(stitchingFile)
	if err != nil {
		log.Println("Error creating stitchingFile: ", err)
	}
	clips.GetTop("itmejp", "2", "all")
	for _, clip := range clips.Clips {
		err = clip.Download()
		if err != nil {
			t.Error(err)
		}
		err = clip.Prepare(stitchingFile)
		if err != nil {
			t.Error(err)
		}
	}
	clips.Stitch(outputFile, stitchingFile)
	if _, err := os.Stat(a.Config.Path + "/stitched.mp4"); err != nil {
		t.Error("Did not stitch")
	}
	log.Println("Starting TestStitch test 2")
	clips.Stitch(outputFile, stitchingFile)
	if _, err := os.Stat(a.Config.Path + "/stitched.mp4"); err != nil {
		t.Error("Did not stitch")
	}
}
