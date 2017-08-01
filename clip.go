package main

import (
	"log"
	"gopkg.in/resty.v0"
	"encoding/json"
	"strings"
	"os"
	"bytes"
	"io"
	"time"
	"github.com/go-redis/redis"
)

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

// Get gets information about Clips from Twitch
func (clip *Clip) Get() error {
	if clip.Slug == "" {
		err := TwitchError("Clip should have a slug")
		log.Printf("Error on Get clip: %s\n", err)
		return err
	}

	log.Printf("Getting information about %s from Twitch", clip.Slug)
	resp, err := resty.R().
		SetHeader("Client-ID", a.Config.ClientID).
		SetHeader("Accept", "application/vnd.twitchtv.v5+json").
		Get(urlClip + clip.Slug)
	if err != nil {
		log.Printf("Error getting clip from Twitch: %s\n", err)
	}
	err = json.Unmarshal(resp.Body(), clip)
	if err != nil {
		log.Printf("Error unserializing json: %s\n", err)
		return err
	}
	return nil
}

// Download downloads the clip from Twitch
func (clip *Clip) Download() error {
	videoURL := clip.Thumbnails.Medium
	videoURL = strings.Replace(videoURL, "-preview-480x272.jpg", ".mp4", -1)

	outString := a.Config.Path + "/" + clip.Slug + ".mp4"

	if _, err := os.Stat(a.Config.Path); os.IsNotExist(err) {
		err := os.Mkdir(a.Config.Path, 0777)
		if err != nil {
			log.Printf("Error creating folder: %s\n", err)
			return err
		}
	}

	out, err := os.Create(outString)
	defer out.Close()
	if err != nil {
		log.Printf("Error creating %s: %s\n", outString, err)
		return err
	}

	err = os.Chmod(outString, 0755)
	if err != nil {
		log.Printf("Error assigning permissions to file %s: %s\n", outString, err)
		return err
	}

	log.Printf("Downloading %s...\n", clip.Slug)
	resp, err := resty.R().
		Get(videoURL)
	if err != nil {
		log.Println("Error getting video data: ", err)
		return err
	}

	r := bytes.NewReader(resp.Body())

	_, err = io.Copy(out, r)
	if err != nil {
		log.Printf("Error saving video data: %s", err)
		return err
	}
	return nil
}