package server

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"microblog/infrastructure/adapters/cache"
	"microblog/infrastructure/adapters/persistence"
)

// NewHTTPServer initializes the HTTP server with all routes and dependencies.
func NewHTTPServer(mongoClient *mongo.Client, redisCache *cache.RedisCache) *gin.Engine {
	r := gin.Default()
	database := mongoClient.Database("microblog")
	userTimelineCollection := database.Collection("user_timeline")
	tweetsCollection := database.Collection("tweets")

	mongoUserTimelineRepository := persistence.NewMongoUserTimelineRepository(userTimelineCollection)
	mongoTweetRepository := persistence.NewMongoTweetRepository(tweetsCollection)
	RegisterRoutes(r, mongoUserTimelineRepository, mongoTweetRepository, redisCache)
	return r
}
