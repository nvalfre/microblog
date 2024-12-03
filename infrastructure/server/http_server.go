package server

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"microblog/infrastructure/adapters/cache"
)

// NewHTTPServer initializes the HTTP server with all routes and dependencies.
func NewHTTPServer(mongoClient *mongo.Client, redisCache *cache.RedisCache) *gin.Engine {
	r := gin.Default()
	RegisterRoutes(r, mongoClient, redisCache)
	return r
}
