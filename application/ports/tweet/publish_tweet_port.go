package tweet

import "microblog/domain/models"

type PublishTweetPort interface {
	Execute(tweet *models.Tweet) error
}
