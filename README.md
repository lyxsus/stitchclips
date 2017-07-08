# stitchclips [![Build Status](https://travis-ci.org/Sadzeih/stitchclips.svg?branch=master)](https://travis-ci.org/Sadzeih/stitchclips)

This API stiches your clips from Twitch into one video automatically.

### Documentation

Visit the [wiki for documentation on the API](https://github.com/Sadzeih/stitchclips/wiki)

## Installation

### Required packages

* ffmpeg
* redis

First download dependencies and build the project

```bash
go get gopkg.in/resty.v0
go get github.com/gorilla/mux
go get github.com/satori/go.uuid
go build
```

or alternatively

```bash
go get github.com/Sadzeih/stitchclips
```

Then you need to create a config file that suits your needs, like so:
```json
{
	"clientId": "TWITCH CLIENT ID",
	"host": "http://localhost",
	"port": "8000",
	"path": "clips_test",
	"redis": {
		"host": "localhost",
		"port": "6379",
		"passsword": "",
		"db": 0
	}
}
```
*Note: you need to have one config file per environment. If you want a config file for production then name it `production`*

## Usage

If you want to run stitchclips with the `production` config file, use:

```bash
GOENV=production ./stitchclips
```
*Note: if you don't precise the GOENV, it will default to `dev`*

## Tests

To run tests, precise the config file you want, for example:

```bash
GOENV=test go test
```

## Known issues

* Simultaneous downloading/encoding:
	* ~~ffmpeg only allows one concatenation at a time it seems~~
	* [downloading a file must be done once to avoid writing problems](https://github.com/Sadzeih/stitchclips/issues/1)
