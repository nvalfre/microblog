package dto

// PublishTweetRequest represents a request to publish a tweet.
type PublishTweetRequest struct {
	UserID  string `json:"user_id" binding:"required"`
	Content string `json:"content" binding:"required"`
}
