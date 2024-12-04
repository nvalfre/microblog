package database

import (
	"context"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
)

func TestNewRedisClient(t *testing.T) {
	mockRedis := miniredis.RunT(t)
	defer mockRedis.Close()

	type args struct {
		addr string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Valid Redis Address",
			args: args{
				addr: mockRedis.Addr(),
			},
			wantErr: false,
		},
		{
			name: "Invalid Redis Address",
			args: args{
				addr: "invalid:6379",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var client *redis.Client
			defer func() {
				if client != nil {
					client.Close()
				}
			}()

			defer func() {
				if r := recover(); r != nil && !tt.wantErr {
					t.Errorf("Unexpected panic for address %s: %v", tt.args.addr, r)
				} else if r == nil && tt.wantErr {
					t.Errorf("Expected error for address %s, but got no panic", tt.args.addr)
				}
			}()

			client = NewRedisClient(tt.args.addr)

			if !tt.wantErr && client != nil {
				err := client.Ping(context.Background()).Err()
				if err != nil {
					t.Errorf("Failed to ping Redis: %v", err)
				}
			}
		})
	}
}
