package models

import (
	"reflect"
	"testing"
	"time"
)

func TestNewTweet(t *testing.T) {
	const characters = 281
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
		{
			name: "Valid tweet creation",
			args: args{
				userID:  "user123",
				content: "This is a valid tweet",
			},
			want: &Tweet{
				UserID:  "user123",
				Content: "This is a valid tweet",
			},
			wantErr: false,
		},
		{
			name: "Tweet content exceeds 280 characters",
			args: args{
				userID:  "user123",
				content: string(make([]byte, characters)),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Empty content",
			args: args{
				userID:  "user123",
				content: "",
			},
			want: &Tweet{
				UserID:  "user123",
				Content: "",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTweet(tt.args.userID, tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTweet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil {
				got.ID = ""
				got.CreatedAt = time.Time{}
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTweet() got = %v, want %v", got, tt.want)
			}
		})
	}
}
