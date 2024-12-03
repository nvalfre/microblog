package tweet

import (
	"microblog/domain/models"
	"microblog/services"
)

// PublishTweetUseCase handles the logic of publishing a new tweet.
type PublishTweetUseCase struct {
	TweetService services.TweetService
}

// Execute publishes a new tweet by saving it to the repository.
func (uc *PublishTweetUseCase) Execute(tweet *models.Tweet) error {
	return uc.TweetService.SaveTweet(tweet)
}
