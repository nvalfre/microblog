package database

import (
	"context"
	"testing"
	"time"
)

func TestNewMongoClient(t *testing.T) {
	type args struct {
		uri string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Valid MongoDB URI",
			args: args{
				uri: "mongodb://localhost:27017",
			},
			wantErr: false,
		},
		{
			name: "Invalid MongoDB URI",
			args: args{
				uri: "invalid-uri",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil && !tt.wantErr {
					t.Errorf("Unexpected panic for URI %s: %v", tt.args.uri, r)
				}
			}()

			got := NewMongoClient(tt.args.uri)
			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			defer cancel()

			err := got.Ping(ctx, nil)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewMongoClient() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
