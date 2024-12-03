package controllers

import (
	"microblog/application/ports/tweet"
	"microblog/domain/models"
	"microblog/interface/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TweetController handles tweet-related requests.
type TweetController struct {
	PublishTweetUseCase tweet.PublishTweetPort
}

// PublishTweet publishes a new tweet.
func (tc *TweetController) PublishTweet(c *gin.Context) {
	var request dto.PublishTweetRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tweet, err := models.NewTweet(request.UserID, request.Content)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = tc.PublishTweetUseCase.Execute(tweet)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to publish tweet"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Tweet published successfully"})
}
