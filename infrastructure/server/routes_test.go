package server

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"microblog/infrastructure/adapters/cache"
	"testing"
)

func TestRegisterRoutes(t *testing.T) {
	type args struct {
		router      *gin.Engine
		mongoClient *mongo.Client
		redisCache  *cache.RedisCache
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RegisterRoutes(tt.args.router, tt.args.mongoClient, tt.args.redisCache)
		})
	}
}
