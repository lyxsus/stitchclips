package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)

	return rr
}

func TestHandleGetClips(t *testing.T) {
	request, err := http.NewRequest("GET", "/clips/itmejp/all/2", nil)
	if err != nil {
		t.Error("Error on request", err)
	}
	r := executeRequest(request)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Error("Error on reading Body", err)
	}
	clips := Clips{}
	err = json.Unmarshal(body, &clips)
	if err != nil {
		t.Error("Error on unserializing json", err)
	}
	if len(clips.Clips) != 2 {
		t.Error("Should return 2 clips")
	}
}

var testClip2 = Clip{
	Slug: "RockyMoldyCougarMcaT",
	Thumbnails: Thumbnails{
		Medium: "https://clips-media-assets.twitch.tv/23732235232-offset-5096.993999999978-30-preview-480x272.jpg",
	},
}

func TestHandleStitch(t *testing.T) {
	clips := Clips{}
	clips.Clips = append(clips.Clips, testClip)
	clips.Clips = append(clips.Clips, testClip2)
	jsonBody, err := json.Marshal(clips)
	if err != nil {
		t.Error("Error on serializing json", err)
	}
	request, err := http.NewRequest("POST", "/stitch", bytes.NewReader(jsonBody))
	if err != nil {
		t.Error("Error on request", err)
	}
	r := executeRequest(request)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Error("Error on reading Body", err)
	}
	video := StitchedVideo{}
	err = json.Unmarshal(body, &video)
	if err != nil {
		t.Error("Error on unserializing json", err)
	}
	request, err = http.NewRequest("GET", "/video/"+video.ID+".mp4", nil)
	r = executeRequest(request)
	if err != nil {
		t.Error("Error getting video data: ", err)
	}
	if r.Code >= 300 {
		fmt.Printf("%s", r.Body)
		t.Error("Fail stitching")
	}

	request, err = http.NewRequest("POST", "/stitch", bytes.NewReader(jsonBody))
	if err != nil {
		t.Error("Error on request", err)
	}
	r = executeRequest(request)
	body, err = ioutil.ReadAll(r.Body)
	if err != nil {
		t.Error("Error on reading Body", err)
	}
	video = StitchedVideo{}
	err = json.Unmarshal(body, &video)
	if err != nil {
		t.Error("Error on unserializing json", err)
	}
	request, err = http.NewRequest("GET", "/video/"+video.ID+".mp4", nil)
	r = executeRequest(request)
	if err != nil {
		t.Error("Error getting video data: ", err)
	}
	if r.Code >= 300 {
		fmt.Printf("%s", r.Body)
		t.Error("Fail stitching")
	}
}
