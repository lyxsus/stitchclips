package main

import (
	"gopkg.in/resty.v0"
	"encoding/json"
	"time"
	"os"
	"fmt"
	"io"
	"strings"
	"bytes"
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

// Clip represents a clip (short video) on twitch
type Clip struct {
	Slug string `json:"slug"`
	TrackingID string `json:"tracking_id"`
	URL string `json:"url"`
	EmbedURL string `json:"embed_url"`
	EmbedHTML string `json:"embed_html"`
	Broadcaster Broadcaster `json:"broadcaster"`
	Curator Broadcaster `json:"curator"`
	Vod struct {
		ID string `json:"id"`
		URL string `json:"url"`
	} `json:"vod"`
	Game string `json:"game"`
	Language string `json:"language"`
	Title string `json:"title"`
	Views int `json:"views"`
	Duration float64 `json:"duration"`
	CreatedAt time.Time `json:"created_at"`
	Thumbnails struct {
		Medium string `json:"medium"`
		Small string `json:"small"`
		Tiny string `json:"tiny"`
	} `json:"thumbnails"`
}

// Clips represents multiple clips
type Clips struct {
	Clips []Clip`json:"clips"`
	Cursor string `json:"_cursor"`
}

var url = "https://api.twitch.tv/kraken/clips/top"
var clientID = "da2bjk7gdn4zt04ssoq0e7zixvaa3f";

// GetTopClips gets multiple clips from a twitch channel
func (clips *Clips) GetTopClips(channel string, limit string, period string) error {
	resp, err := resty.R().
		SetQueryParams(map[string]string{
			"channel": channel,
			"limit":   limit,
			"period":  period}).
		SetHeader("Client-ID", clientID).
		SetHeader("Accept", "application/vnd.twitchtv.v5+json").
		Get(url)
	if err != nil {
		return nil
	}
	json.Unmarshal(resp.Body(), clips)
	return nil
}

// DownloadClip downloads the clip from Twitch
func (clip *Clip) DownloadClip(dir string) error {
	videoURL := clip.Thumbnails.Medium
	videoURL = strings.Replace(videoURL, "-preview-480x272.jpg", ".mp4", -1)

	outString := dir + "/" + clip.Slug + ".mp4"

	os.Mkdir(dir, 0777)
	
	out, err := os.Create(outString)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := resty.R().
		Get(videoURL)
		
	r := bytes.NewReader(resp.Body())
	
	io.Copy(out, r)
	if err != nil {
		return err
	}
	return nil
}
