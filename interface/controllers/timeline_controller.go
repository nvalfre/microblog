package controllers

import (
	"microblog/application/ports/tweet"
	"microblog/interface/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TimelineController handles timeline-related requests.
type TimelineController struct {
	GetTimelineUseCase tweet.GetTimelinePort
}

// GetTimeline retrieves the timeline of a user.
func (tc *TimelineController) GetTimeline(c *gin.Context) {
	userID := c.Query("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing user_id"})
		return
	}

	tweets, err := tc.GetTimelineUseCase.Execute(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch timeline", "message": err.Error()})
		return
	}

	response := dto.TimelineResponse{Tweets: tweets}
	c.JSON(http.StatusOK, response)
}
