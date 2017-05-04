package main

var config = LoadConfig()

func main() {
	clips := Clips{}

	clips.GetTopClips(config.Channel, config.Limit, config.Period)
	for _, clip := range clips.Clips {
		clip.DownloadClip(config.Path)
	}
	clips.StitchClips(config.Path)
}
