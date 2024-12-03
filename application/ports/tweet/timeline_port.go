package tweet

import "microblog/domain/models"

type GetTimelinePort interface {
	Execute(userID string) ([]models.Tweet, error)
}
