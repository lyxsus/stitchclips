package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"gopkg.in/resty.v0"
)

// TwitchError represents an error string
type TwitchError string

func (err TwitchError) Error() string {
	return fmt.Sprintf("Twitch Error: %s", string(err))
}

// Broadcaster represents a twitch user
type Broadcaster struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	ChannelURL  string `json:"channel_url"`
	Logo        string `json:"logo"`
}

// Thumbnails represents clips thumbnails
type Thumbnails struct {
	Medium string `json:"medium"`
	Small  string `json:"small"`
	Tiny   string `json:"tiny"`
}

// Clip represents a clip (short video) on twitch
type Clip struct {
	Slug        string      `json:"slug"`
	TrackingID  string      `json:"tracking_id"`
	URL         string      `json:"url"`
	EmbedURL    string      `json:"embed_url"`
	EmbedHTML   string      `json:"embed_html"`
	Broadcaster Broadcaster `json:"broadcaster"`
	Curator     Broadcaster `json:"curator"`
	Vod         struct {
		ID  string `json:"id"`
		URL string `json:"url"`
	} `json:"vod"`
	Game       string     `json:"game"`
	Language   string     `json:"language"`
	Title      string     `json:"title"`
	Views      int        `json:"views"`
	Duration   float64    `json:"duration"`
	CreatedAt  time.Time  `json:"created_at"`
	Thumbnails Thumbnails `json:"thumbnails"`
}

// Clips represents multiple clips
type Clips struct {
	Clips  []Clip `json:"clips"`
	Cursor string `json:"_cursor"`
}

var url = "https://api.twitch.tv/kraken/clips/top"

// GetTop gets multiple clips from a twitch channel
func (clips *Clips) GetTop(channel string, limit string, period string) {
	log.Printf("Getting the top %s clips for %s during the last %s", limit, channel, period)
	resp, err := resty.R().
		SetQueryParams(map[string]string{
			"channel": channel,
			"limit":   limit,
			"period":  period}).
		SetHeader("Client-ID", config.ClientID).
		SetHeader("Accept", "application/vnd.twitchtv.v5+json").
		Get(url)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(resp.Body(), clips)
	if err != nil {
		log.Fatal(err)
	}
}

// Download downloads the clip from Twitch
func (clip *Clip) Download() {
	videoURL := clip.Thumbnails.Medium
	videoURL = strings.Replace(videoURL, "-preview-480x272.jpg", ".mp4", -1)

	outString := config.Path + "/" + clip.Slug + ".mp4"

	if _, err := os.Stat(config.Path); os.IsNotExist(err) {
		err := os.Mkdir(config.Path, 0777)
		if err != nil {
			log.Fatal(err)
		}
	}

	out, err := os.Create(outString)
	defer out.Close()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Downloading %s...\n", clip.Slug)
	resp, err := resty.R().
		Get(videoURL)
	if err != nil {
		log.Fatal(err)
	}

	r := bytes.NewReader(resp.Body())

	_, err = io.Copy(out, r)
	if err != nil {
		log.Fatal(err)
	}
}
