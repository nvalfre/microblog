package repository

import "microblog/domain/models"

type TweetRepository interface {
	Save(tweet *models.Tweet) error
	GetTimeline(userIDs []string) ([]models.Tweet, error)
}
