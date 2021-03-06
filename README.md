 [![Build Status](https://travis-ci.org/Sadzeih/stitchclips.svg?branch=master)](https://travis-ci.org/Sadzeih/stitchclips) [![Codacy Badge](https://api.codacy.com/project/badge/Grade/d03cd0dedcb14f0f928a09bdf72fd219)](https://www.codacy.com/app/Stitchclips/stitchclips?utm_source=github.com&utm_medium=referral&utm_content=Sadzeih/stitchclips&utm_campaign=badger) 
# stitchclips

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
go get github.com/rs/cors
go get github.com/spf13/viper
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
	"path": "clips_test"
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
