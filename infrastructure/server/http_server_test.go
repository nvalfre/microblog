package server

import (
	"microblog/infrastructure/adapters/cache"
	"testing"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestNewHTTPServer(t *testing.T) {
	type args struct {
		mongoClient *mongo.Client
		redisCache  *cache.RedisCache
	}
	tests := []struct {
		name string
		args args
		want *gin.Engine
	}{
		{
			name: "Valid dependencies",
			args: args{
				mongoClient: mockMongoClient(),
				redisCache:  mockRedisCache(),
			},
			want: mockGinEngine(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewHTTPServer(tt.args.mongoClient, tt.args.redisCache)
			if !compareGinEngines(got, tt.want) {
				t.Errorf("NewHTTPServer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func mockMongoClient() *mongo.Client {
	client, _ := mongo.NewClient(options.Client().ApplyURI("mongodb://mockhost:27017"))
	return client
}

func mockRedisCache() *cache.RedisCache {
	return &cache.RedisCache{}
}

func mockGinEngine() *gin.Engine {
	r := gin.Default()
	return r
}

func compareGinEngines(got, want *gin.Engine) bool {
	return got != nil && want != nil
}
