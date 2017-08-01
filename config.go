package main

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

// Config represents stitchclips configuration
type Config struct {
	ClientID string
	Host     string
	Port     string
	Path     string
	DBURL    string
	DBName   string
}

// LoadConfig loads config from file
// You need to set GOENV to the program's environment
// example: GOENV=test or GOENV=prod
// You need to have a config file that matches the environment provided
func LoadConfig() {
	viper.SetDefault("host", "localhost")
	viper.SetDefault("port", "8080")
	viper.SetDefault("path", "clips")
	viper.SetDefault("mongodb_url", "mongodb://localhost:27017")
	viper.SetDefault("db", "stitchclips")

	env := os.Getenv("GOENV")
	if env == "" {
		env = "dev"
	}
	viper.SetConfigName(env)
	viper.AddConfigPath("config/")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Fatal error config file: ", err)
	}

	if viper.IsSet("clientId") {
		a.Config.ClientID = viper.GetString("clientId")
	} else {
		log.Fatal("No client ID was set in config file.")
	}
	a.Config.Host = viper.GetString("host")
	a.Config.Port = viper.GetString("port")
	a.Config.Path = viper.GetString("path")
	a.Config.DBURL = viper.GetString("mongodb_url")
	a.Config.DBName = viper.GetString("db")
}
