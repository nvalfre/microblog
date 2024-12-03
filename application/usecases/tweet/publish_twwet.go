package tweet

import (
	"microblog/domain/models"
	"microblog/domain/repository"
)

// PublishTweetUseCase handles the logic of publishing a new tweet.
type PublishTweetUseCase struct {
	TweetRepo repository.TweetRepository
}

// Execute publishes a new tweet by saving it to the repository.
func (uc *PublishTweetUseCase) Execute(tweet *models.Tweet) error {
	return uc.TweetRepo.Save(tweet)
}
