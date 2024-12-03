package main

import (
	"microblog/infrastructure/adapters/cache"
	"microblog/infrastructure/database"
	"microblog/infrastructure/logger"
	"microblog/infrastructure/server"
)

func main() {
	logger.InitializeLogger()
	mongoClient := database.NewMongoClient("mongodb://localhost:27017")
	redisClient := database.NewRedisClient("localhost:6379")
	redisCache := cache.NewRedisCache(redisClient)

	srv := server.NewHTTPServer(mongoClient, redisCache)
	if err := srv.Run(":8080"); err != nil {
		logger.Fatal("Failed to start server", err)
	}
}
