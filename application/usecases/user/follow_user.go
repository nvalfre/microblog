package user

import (
	"microblog/services"
)

// FollowUserUseCase handles the logic of follow users.
type FollowUserUseCase struct {
	UserService services.UserService
}

// Execute follow target user id for a user
func (uc *FollowUserUseCase) Execute(userID, targetID string) error {
	return uc.UserService.AddFollower(userID, targetID)
}
