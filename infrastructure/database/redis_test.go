package database

import (
	"github.com/redis/go-redis/v9"
	"reflect"
	"testing"
)

func TestNewRedisClient(t *testing.T) {
	type args struct {
		addr string
	}
	tests := []struct {
		name string
		args args
		want *redis.Client
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRedisClient(tt.args.addr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRedisClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
