package main

import (
	"encoding/json"
	"fmt"
	"log"
	"gopkg.in/resty.v0"
)

// TwitchError represents an error string
type TwitchError string

func (err TwitchError) Error() string {
	return fmt.Sprintf("Twitch Error: %s", string(err))
}

// Clips represents multiple clips
type Clips struct {
	Clips  []Clip `json:"clips"`
	Cursor string `json:"_cursor"`
	ID     string
}

const urlTop = "https://api.twitch.tv/kraken/clips/top"
const urlClip = "https://api.twitch.tv/kraken/clips/"

// GetTop gets multiple clips from a twitch channel
func (clips *Clips) GetTop(channel string, limit string, period string) error {
	log.Printf("Getting the top %s clips for %s during the last %s", limit, channel, period)
	resp, err := resty.R().
		SetQueryParams(map[string]string{
			"channel": channel,
			"limit":   limit,
			"period":  period}).
		SetHeader("Client-ID", a.Config.ClientID).
		SetHeader("Accept", "application/vnd.twitchtv.v5+json").
		Get(urlTop)
	if err != nil {
		log.Printf("Error getting clips from Twitch: %s\n", err)
		return err
	}
	err = json.Unmarshal(resp.Body(), clips)
	if err != nil {
		log.Printf("Error unserializing json: %s\n", err)
		return err
	}
	return nil
}

// Slugs returns the slugs of all clips concatenated with a space
func (clips *Clips) Slugs() string {
	str := ""
	for index, clip := range clips.Clips {
		if index != 0 {
			str += " "
		}
		str += clip.Slug
	}
	return str
}

func (clips *Clips) DownloadAll() error {
	clipMap := make(map[string]Clip)
	for _, clip := range clips.Clips {
		clipMap[clip.Slug] = clip
	}
	out, error := a.Dm.addClips(*clips)
	for len(clipMap) != 0 {
		slug := <-out
		err := <-error
		if err != nil {
			return err
		}
		delete(clipMap, slug)
	}
	return nil
}

