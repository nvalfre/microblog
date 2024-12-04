package main

import (
	"github.com/sirupsen/logrus"
	"microblog/config"
	"microblog/infrastructure/adapters/cache"
	"microblog/infrastructure/database"
	"microblog/infrastructure/logger"
	"microblog/infrastructure/server"
)

func main() {
	logger.InitializeLogger()

	appConfig, err := config.LoadConfig("./config")

	logger.Info("Connecting to MongoDB at URI", logrus.Fields{"appConfig": appConfig})

	if err != nil {
		logger.Fatal("Failed to load configuration", err)
	}

	mongoClient := database.NewMongoClient(appConfig.Database.URI)
	redisClient := database.NewRedisClient(appConfig.Redis.Address)
	redisCache := cache.NewRedisCache(redisClient)

	srv := server.NewHTTPServer(mongoClient, redisCache)
	if err := srv.Run(":8080"); err != nil {
		logger.Fatal("Failed to start server", err)
	}
}
