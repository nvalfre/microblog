package user

import "microblog/domain/repository"

// FollowUserUseCase handles the logic of follow users.
type FollowUserUseCase struct {
	UserRepo repository.UserRepository
}

// Execute follow target user id for a user
func (uc *FollowUserUseCase) Execute(userID, targetID string) error {
	return uc.UserRepo.AddFollower(userID, targetID)
}
