package main

import (
	"fmt"
)

func main() {
	clips := Clips{}

	err := clips.GetTopClips("itmejp", "10", "week")
	if err != nil {
		fmt.Println(err)
	}
	for _, clip := range clips.Clips {
		err = clip.DownloadClip("clips")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Clip:\n\tslug: %s\n\turl: %s\n\tauthor: %s\n", clip.Slug, clip.EmbedURL, clip.Curator.Name)
	}
}
