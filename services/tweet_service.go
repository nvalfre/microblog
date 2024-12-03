package services

import (
	"microblog/domain/models"
	"microblog/domain/repository"
)

type TweetService interface {
	SaveTweet(tweet *models.Tweet) error
	GetTimeline(userIDs []string) ([]models.Tweet, error)
}

// TweetServiceImpl provides implementation methods to interact with the tweet repository.
type TweetServiceImpl struct {
	TweetRepo repository.TweetRepository
}

// SaveTweet saves a tweet to the database.
func (s *TweetServiceImpl) SaveTweet(tweet *models.Tweet) error {
	return s.TweetRepo.Save(tweet)
}

// GetTimeline retrieves the timeline of tweets for a given set of user IDs.
func (s *TweetServiceImpl) GetTimeline(userIDs []string) ([]models.Tweet, error) {
	return s.TweetRepo.GetTimeline(userIDs)
}

// NewTweetService creates a new instance of TweetService.
func NewTweetService(tweetRepo repository.TweetRepository) TweetService {
	return &TweetServiceImpl{TweetRepo: tweetRepo}
}
