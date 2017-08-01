package main

import (
	"bytes"
	"encoding/json"
	"gopkg.in/resty.v0"
	"io"
	"log"
	"os"
	"strings"
	"time"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
)

// Broadcaster represents a twitch user
type Broadcaster struct {
	ID          string `bson:"id" json:"id"`
	Name        string `bson:"name" json:"name"`
	DisplayName string `bson:"display_name" json:"display_name"`
	ChannelURL  string `bson:"channel_url" json:"channel_url"`
	Logo        string `bson:"logo" json:"logo"`
}

// Thumbnails represents clips thumbnails
type Thumbnails struct {
	Medium string `bson:"medium" json:"medium"`
	Small  string `bson:"small" json:"small"`
	Tiny   string `bson:"tiny" json:"tiny"`
}

// Clip represents a clip (short video) on twitch
type Clip struct {
	Slug        string      `bson:"slug" json:"slug"`
	TrackingID  string      `bson:"tracking_id" json:"tracking_id"`
	URL         string      `bson:"url" json:"url"`
	EmbedURL    string      `bson:"embed_url" json:"embed_url"`
	EmbedHTML   string      `bson:"embed_html" json:"embed_html"`
	Broadcaster Broadcaster `bson:"broadcaster" json:"broadcaster"`
	Curator     Broadcaster `bson:"curator" json:"curator"`
	Vod         struct {
		ID  string `bson:"id" json:"id"`
		URL string `bson:"url" json:"url"`
	} `bson:"vod" json:"vod"`
	Game       string     `bson:"game" json:"game"`
	Language   string     `bson:"language" json:"language"`
	Title      string     `bson:"title" json:"title"`
	Views      int        `bson:"views" json:"views"`
	Duration   float64    `bson:"duration" json:"duration"`
	CreatedAt  time.Time  `bson:"created_at" json:"created_at"`
	Thumbnails Thumbnails `bson:"thumbnails" json:"thumbnails"`
}

// GetFromTwitch gets information about Clips from Twitch
func (clip *Clip) GetFromTwitch() error {
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

	a.Db.C("clips").Insert(clip)

	index := mgo.Index{
		Key: []string{"slug"},
		Unique: true,
		DropDups: true,
		Background: false,
		Sparse: true,
	}
	err = a.Db.C("clips").EnsureIndex(index)
	if err != nil {
		log.Println("Could not ensure index Slug existed", err)
	}

	return nil
}

func (clip *Clip) Get() error {
	if clip.Slug == "" {
		err := TwitchError("Clip should have a slug")
		log.Printf("Error on Get clip: %s\n", err)
		return err
	}

	searchedClip := []Clip{}
	err := a.Db.C("clips").Find(bson.M{"slug": clip.Slug}).Limit(1).All(&searchedClip)
	if err != nil {
		log.Println(err)
		return err
	}

	if len(searchedClip) == 0 {
		return clip.GetFromTwitch()
	}
	clip = &searchedClip[0]

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