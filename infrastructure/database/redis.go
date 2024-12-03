package database

import (
	"context"
	"github.com/redis/go-redis/v9"
	"microblog/infrastructure/logger"
)

// NewRedisClient initializes a new Redis client.
func NewRedisClient(addr string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		logger.Fatal("Failed to connect to Redis", err)
	}
	return client
}
