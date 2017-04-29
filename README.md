# stitchclips

This tool stiches your clips from Twitch automatically and post them to Youtube or Twitch Videos

### Features

* Downloads top clips from a Twitch channel
	* Choose how many clips you want
	* Choose the channel
	* Choose the period: last day, week, month, all
	* Choose where you want them saved

## Installation

First download dependencies and build the project

```bash
go get gopkg.in/resty.v0
go build
```

Then you need to create a config file that suits your needs, like so:
```json
{
	"clientId": "TWITCH CLIENT ID",
	"channel": "itmejp",
	"period": "week",
	"limit": "10",
	"path": "clips"
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