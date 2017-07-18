package main

import (
	"log"
	"os"
	"github.com/spf13/viper"
)

// RedisConfig contains needed configuration to connect to the Redis DB
type RedisConfig struct {
	Host     string
	Port     string
	Password string
	Db       int
}

// Config represents stitchclips configuration
type Config struct {
	ClientID string
	Host     string
	Port     string
	Path     string
	Redis    RedisConfig
}

// LoadConfig loads config from file
// You need to set GOENV to the program's environment
// example: GOENV=test or GOENV=prod
// You need to have a config file that matches the environment provided
func LoadConfig() {
	viper.SetDefault("host", "localhost")
	viper.SetDefault("port", "8080")
	viper.SetDefault("path", "clips")
	viper.SetDefault("redis.host", "localhost")
	viper.SetDefault("redis.port", "6379")
	viper.SetDefault("redis.password", "")
	viper.SetDefault("redis.db", 0)
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
	a.Config.Redis.Host = viper.GetString("redis.host")
	a.Config.Redis.Port = viper.GetString("redis.port")
	a.Config.Redis.Password = viper.GetString("redis.password")
	a.Config.Redis.Db = viper.GetInt("redis.db")
}
