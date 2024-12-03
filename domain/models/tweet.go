package models

import (
	"errors"
	"time"
)

type Tweet struct {
	ID        string    `json:"id" bson:"id"`
	UserID    string    `json:"user_id" bson:"user_id"`
	Content   string    `json:"content" bson:"content"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}

func NewTweet(userID, content string) (*Tweet, error) {
	if len(content) > 280 {
		return nil, errors.New("content exceeds 280 characters")
	}
	return &Tweet{
		ID:        generateID(),
		UserID:    userID,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func generateID() string {
	return time.Now().Format("20060102150405")
}
