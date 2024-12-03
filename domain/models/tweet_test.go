package models

import (
	"reflect"
	"testing"
)

func TestNewTweet(t *testing.T) {
	type args struct {
		userID  string
		content string
	}
	tests := []struct {
		name    string
		args    args
		want    *Tweet
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTweet(tt.args.userID, tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTweet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTweet() got = %v, want %v", got, tt.want)
			}
		})
	}
}
