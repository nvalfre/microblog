package main

import (
	"log"
	"microblog/config"
	"microblog/infrastructure/adapters/cache"
	"microblog/infrastructure/database"
	"microblog/infrastructure/logger"
	"microblog/infrastructure/server"
	"os"
)

func main() {
	logger.InitializeLogger()

	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}

	appConfig, err := config.LoadConfig("./config", env)
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	mongoClient := database.NewMongoClient(appConfig.Database.URI)
	redisClient := database.NewRedisClient(appConfig.Redis.Address)
	redisCache := cache.NewRedisCache(redisClient)

	srv := server.NewHTTPServer(mongoClient, redisCache)
	if err := srv.Run(":8080"); err != nil {
		logger.Fatal("Failed to start server", err)
	}
}
