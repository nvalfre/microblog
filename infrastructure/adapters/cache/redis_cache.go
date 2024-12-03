package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

// RedisCache is a wrapper around a Redis client for caching.
type RedisCache struct {
	Client *redis.Client
}

// NewRedisCache initializes a new RedisCache.
func NewRedisCache(client *redis.Client) *RedisCache {
	return &RedisCache{Client: client}
}

// Get retrieves a value from the cache.
func (c *RedisCache) Get(key string) (string, error) {
	return c.Client.Get(context.Background(), key).Result()
}

// Set stores a value in the cache with an expiration.
func (c *RedisCache) Set(key string, value string, expiration int) error {
	return c.Client.Set(context.Background(), key, value, time.Duration(expiration)*time.Second).Err()
}

// Delete removes a value from the cache.
func (c *RedisCache) Delete(key string) error {
	return c.Client.Del(context.Background(), key).Err()
}
