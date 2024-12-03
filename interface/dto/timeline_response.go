package dto

import "microblog/domain/models"

// TimelineResponse represents a user's timeline response.
type TimelineResponse struct {
	Tweets []models.Tweet `json:"tweets"`
}
