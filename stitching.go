package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

// StitchClips uses ffmpeg to concatenate clips .mp4 videos together into stitched.mp4
func (clips Clips) StitchClips(path string) {
	if _, err := os.Stat("stitched.mp4"); err == nil {
		err = os.Remove("stitched.mp4")
		if err != nil {
			log.Fatal(err)
		}
	}

	file, err := os.Create("stitching")
	if err != nil {
		log.Fatal(err)
	}

	for _, clip := range clips.Clips {
		mp4Path := path + "/" + clip.Slug + ".mp4"
		mpgPath := path + "/" + clip.Slug + ".mpg"
		if _, err := os.Stat(mpgPath); err == nil {
			err := os.Remove(mpgPath)
			if err != nil {
				log.Fatal(err)
			}
		}
		log.Println("Encoding " + clip.Slug + " to .mpg...")
		cmd := exec.Command("ffmpeg", "-i", mp4Path, "-qscale", "0", mpgPath)
		err = cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
		concatPath := fmt.Sprintf("file '" + mpgPath + "'\n")
		io.WriteString(file, concatPath)
	}
	log.Println("Sitching...")
	cmd := exec.Command("ffmpeg", "-f", "concat", "-i", "stitching", "-qscale", "0", "-vcodec", "mpeg4", "-c", "copy", "stitched.mp4")
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
