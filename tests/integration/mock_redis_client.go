package integration

import (
	"context"
	"github.com/stretchr/testify/mock"
)

// MockRedisClient mocks Redis client behavior.
type MockRedisClient struct {
	mock.Mock
}

// Get retrieves a value from the cache.
func (c *MockRedisClient) Get(key string) (string, error) {
	args := c.Called(context.Background(), key)
	return args.String(0), args.Error(1)
}

// Set stores a value in the cache with an expiration.
func (c *MockRedisClient) Set(key string, value string, expiration int) error {
	args := c.Called(context.Background(), key, value, expiration)
	return args.Error(0)
}

// Delete removes a value from the cache.
func (c *MockRedisClient) Delete(key string) error {
	args := c.Called(context.Background(), key)
	return args.Error(0)
}
