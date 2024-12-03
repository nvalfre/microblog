package database

import (
	"go.mongodb.org/mongo-driver/mongo"
	"reflect"
	"testing"
)

func TestNewMongoClient(t *testing.T) {
	type args struct {
		uri string
	}
	tests := []struct {
		name string
		args args
		want *mongo.Client
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMongoClient(tt.args.uri); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMongoClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
