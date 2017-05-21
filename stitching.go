package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

// ToMPGAsync encodes a MP4 into a MPG asynchronously
func (clip Clip) ToMPGAsync(stichingFile string, done chan bool) error {
	err := clip.toMPGSync(stichingFile)
	done <- true
	return err
}

// ToMPG encodes a MP4 into a MPG synchronously
func (clip Clip) ToMPG(stichingFile string) error {
	err := clip.toMPGSync(stichingFile)
	return err
}

// ToMPGSync encodes a MP4 into a MPG synchronously
func (clip Clip) toMPGSync(stichingFile string) error {
	mpgPath := config.Path + "/" + clip.Slug + ".mpg"

	if _, err := os.Stat(mpgPath); os.IsNotExist(err) {
		log.Println("Encoding " + clip.Slug + " to .mpg...")
		mp4Path := config.Path + "/" + clip.Slug + ".mp4"
		cmd := exec.Command("ffmpeg", "-i", mp4Path, mpgPath)
		err := cmd.Run()
		if err != nil {
			log.Printf("Error on running encoding %s into mpg: %s\n", mp4Path, err)
			return err
		}
		log.Println("Done encoding " + clip.Slug)
	}

	var file *os.File
	log.Printf("Opening stitching file: %s.\n", stichingFile)
	file, err := os.OpenFile(stichingFile, os.O_RDWR|os.O_APPEND, 0777)
	if err != nil {
		log.Printf("Error opening %s: %s\n", stichingFile, err)
		return err
	}
	defer file.Close()
	concatPath := fmt.Sprintf("file '" + mpgPath + "'\n")
	_, err = io.WriteString(file, concatPath)
	if err != nil {
		log.Printf("Error writing into %s: %s\n", stichingFile, err)
		return err
	}
	return nil
}

// Cleanup deletes the .mp4 video
func (clip Clip) Cleanup() error {
	mp4Path := config.Path + "/" + clip.Slug + ".mp4"
	if _, err := os.Stat(mp4Path); err == nil {
		log.Println("Cleaning " + mp4Path)
		err := os.Remove(mp4Path)
		if err != nil {
			log.Printf("Error removing %s: %s\n", mp4Path, err)
			return err
		}
	}
	return nil
}

// Stitch uses ffmpeg to concatenate clips .mp4 videos together into stitched.mp4
func (clips Clips) Stitch(outputFile string, stitchingFile string) error {
	log.Println("Sitching...")
	cmd := exec.Command("ffmpeg", "-f", "concat", "-i", stitchingFile, "-vcodec", "mpeg4", "-c", "copy", outputFile+".mp4")
	// cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Printf("Error running stitching: %s\n", err)
		return err
	}
	return nil
}
