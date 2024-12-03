package server

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"microblog/infrastructure/adapters/cache"
	"reflect"
	"testing"
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHTTPServer(tt.args.mongoClient, tt.args.redisCache); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHTTPServer() = %v, want %v", got, tt.want)
			}
		})
	}
}
