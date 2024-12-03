package tweet

import (
	"encoding/json"
	"errors"
	"microblog/domain/models"
	"microblog/infrastructure/adapters/interfaces"
	"microblog/services"
)

const cacheExpiration = 60

// GetTimelineUseCase handles the logic of get timeline of followed users.
type GetTimelineUseCase struct {
	TweetService services.TweetService
	UserServuce  services.UserService
	Cache        interfaces.Cache
}

// Execute get timeline to fetch tweets for all followed users.
func (uc *GetTimelineUseCase) Execute(userID string) ([]models.Tweet, error) {
	cacheKey := "timeline:" + userID

	cachedTimeline, err := uc.Cache.Get(cacheKey)
	if err == nil && cachedTimeline != "" {
		var tweets []models.Tweet
		if err := json.Unmarshal([]byte(cachedTimeline), &tweets); err == nil {
			return tweets, nil
		}
	}

	followers, err := uc.UserServuce.GetFollowers(userID)
	if err != nil {
		return nil, err
	}

	tweets, err := uc.TweetService.GetTimeline(followers)
	if err != nil {
		return nil, err
	}

	if tweets == nil {
		return nil, errors.New("tweets is empty")
	}

	timelineJSON, _ := json.Marshal(tweets)
	uc.Cache.Set(cacheKey, string(timelineJSON), cacheExpiration)

	return tweets, nil
}
