package main

var config = LoadConfig()

func main() {
	clips := Clips{}

	clips.GetTop(config.Channel, config.Limit, config.Period)

	done := make(chan bool, len(clips.Clips))
	for _, clip := range clips.Clips {
		clip.Download()
		go clip.ToMPGAsync(done)
	}
	for i := 0; i < len(clips.Clips); i++ {
		<-done
	}
	clips.Stitch()
	for _, clip := range clips.Clips {
		clip.Cleanup()
	}
}
