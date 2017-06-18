package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

// PrepareAsync encodes a MP4 into an intermediate asynchronously
func (clip Clip) PrepareAsync(stichingFile string, done chan bool, errors chan error) error {
	err := clip.PrepareSync(stichingFile)
	done <- true
	errors <- err
	return err
}

// Prepare encodes a MP4 into a intermediate file synchronously
func (clip Clip) Prepare(stichingFile string) error {
	err := clip.PrepareSync(stichingFile)
	return err
}

// PrepareSync encodes a MP4 into a intermediate file synchronously
func (clip Clip) PrepareSync(stichingFile string) error {
	intermediatePath := a.Config.Path + "/" + clip.Slug + ".ts"

	if _, err := os.Stat(intermediatePath); os.IsNotExist(err) {
		log.Println("Encoding " + clip.Slug + " to .mpg...")
		mp4Path := a.Config.Path + "/" + clip.Slug + ".mp4"
		cmd := exec.Command("ffmpeg", "-i", mp4Path, "-c", "copy", "-bsf:v", "h264_mp4toannexb", "-f", "mpegts", intermediatePath)
		err := cmd.Run()
		if err != nil {
			log.Printf("Error on running encoding %s into mpg: %s\n", mp4Path, err)
			return err
		}
		log.Println("Done encoding " + clip.Slug)
	}
	return nil
}

// Cleanup deletes the .mp4 video
func (clip Clip) Cleanup() error {
	mp4Path := a.Config.Path + "/" + clip.Slug + ".mp4"
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
	if _, err := os.Stat(outputFile); err == nil {
		log.Println("Already stitched")
		return nil
	}

	log.Println("Sitching...")
	slugs := make([]string, 0)
	for _, clip := range clips.Clips {
		if clip.Slug != "" {
			err := os.Chmod(a.Config.Path+"/"+clip.Slug+".ts", 0755)
			if err != nil {
				log.Printf("Error assigning permissions to file: %s\n", err)
				return err
			}
			slugs = append(slugs, a.Config.Path+"/"+clip.Slug+".ts")
		}
	}
	concatString := strings.Join(slugs, "|")
	cmd := exec.Command("ffmpeg", "-i", "concat:"+concatString, "-c", "copy", "-bsf:a", "aac_adtstoasc", outputFile+".mp4")
	err := cmd.Run()
	if err != nil {
		log.Printf("Error running stitching: %s\n", err)
		return err
	}
	cmd = exec.Command("ffmpeg", "-i", outputFile+".mp4", "-c", "copy", "-movflags", "+faststart", outputFile+"_fs.mp4")
	err = cmd.Run()
	if err != nil {
		log.Printf("Error running stitching: %s\n", err)
		return err
	}
	err = os.Rename(outputFile+"_fs.mp4", outputFile+".mp4")
	if err != nil {
		log.Printf("Error renaming file: %s\n", err)
		return err
	}
	err = os.Chmod(outputFile+".mp4", 0755)
	if err != nil {
		log.Printf("Error assigning permissions to file: %s\n", err)
		return err
	}
	log.Println("Done stitching")
	return nil
}
