package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

// ToMPGAsync encodes a MP4 into a MPG asynchronously
func (clip Clip) ToMPGAsync(done chan bool) {
	clip.toMPGSync()
	done <- true
}

// ToMPG encodes a MP4 into a MPG synchronously
func (clip Clip) ToMPG() {
	clip.toMPGSync()
}

// ToMPGSync encodes a MP4 into a MPG synchronously
func (clip Clip) toMPGSync() {
	mpgPath := config.Path + "/" + clip.Slug + ".mpg"
	mp4Path := config.Path + "/" + clip.Slug + ".mp4"

	if _, err := os.Stat(mpgPath); err == nil {
		err := os.Remove(mpgPath)
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Println("Encoding " + clip.Slug + " to .mpg...")
	cmd := exec.Command("ffmpeg", "-i", mp4Path, mpgPath)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	var file *os.File
	if _, err := os.Stat("stitching"); os.IsNotExist(err) {
		file, err = os.Create("stitching")
		if err != nil {
			log.Fatal(err)
		}
	} else {
		file, err = os.OpenFile("stitching", os.O_RDWR, 0777)
		if err != nil {
			log.Fatal(err)
		}
	}
	defer file.Close()
	concatPath := fmt.Sprintf("file '" + mpgPath + "'\n")
	_, err = io.WriteString(file, concatPath)
	if err != nil {
		log.Fatal(err)
	}
}

// Cleanup cleans the output directory of the MPG file
func (clip Clip) Cleanup() {
	mpgPath := config.Path + "/" + clip.Slug + ".mpg"
	if _, err := os.Stat(mpgPath); err == nil {
		log.Println("Cleaning " + mpgPath)
		err := os.Remove(mpgPath)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// Stitch uses ffmpeg to concatenate clips .mp4 videos together into stitched.mp4
func (clips Clips) Stitch() {
	if _, err := os.Stat(config.Output + ".mp4"); err == nil {
		err = os.Remove(config.Output + ".mp4")
		if err != nil {
			log.Fatal(err)
		}
	}
	log.Println("Sitching...")
	cmd := exec.Command("ffmpeg", "-f", "concat", "-i", "stitching", "-vcodec", "mpeg4", "-c", "copy", config.Output+".mp4")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
