package dto

// FollowUserRequest represents a request to follow a user.
type FollowUserRequest struct {
	UserID   string `json:"user_id" binding:"required"`
	TargetID string `json:"target_id" binding:"required"`
}
