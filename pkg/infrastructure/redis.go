package infrastructure

import (
	redis "github.com/go-redis/redis/v8"
)

type RedisClient struct {
	Client *redis.Client
}

func NewRedis() *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
		DB:   0,
	})
	return &RedisClient{Client: client}
}
