package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// RedisConfig contains needed configuration to connect to the Redis DB
type RedisConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Password string `json:"password"`
	Db       int    `json:"db"`
}

// Config represents stitchclips configuration
type Config struct {
	ClientID string      `json:"clientId"`
	Host     string      `json:"host"`
	Port     string      `json:"port"`
	Path     string      `json:"path"`
	Redis    RedisConfig `json:"redis"`
}

// LoadConfig loads config from file
// You need to set GOENV to the program's environment
// example: GOENV=test or GOENV=prod
// You need to have a config file that matches the environment provided
func LoadConfig() Config {
	config := Config{}
	env := os.Getenv("GOENV")
	if env == "" {
		env = "dev"
	}
	file, err := ioutil.ReadFile("config/" + env + ".json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Fatal(err)
	}
	return config
}
