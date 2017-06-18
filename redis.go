package main

import "github.com/go-redis/redis"

// CreateClient creates a redis client
func CreateClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     a.Config.Redis.Host + ":" + a.Config.Redis.Port,
		Password: a.Config.Redis.Password,
		DB:       a.Config.Redis.Db,
	})

	return client
}
