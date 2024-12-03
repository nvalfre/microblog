package services

import (
	"microblog/domain/repository"
)

type UserService interface {
	AddFollower(followerID, userID string) error
	GetFollowers(userID string) ([]string, error)
}

// UserServiceImpl provides implementation methods to interact with the user repository.
type UserServiceImpl struct {
	UserRepo repository.UserRepository
}

// AddFollower adds a follower to a user.
func (s *UserServiceImpl) AddFollower(userID, followerID string) error {
	return s.UserRepo.AddFollower(userID, followerID)
}

// GetFollowers retrieves all followers of a user.
func (s *UserServiceImpl) GetFollowers(userID string) ([]string, error) {
	return s.UserRepo.GetFollowers(userID)
}

// NewUserService creates a new instance of UserService.
func NewUserService(userRepo repository.UserRepository) UserService {
	return &UserServiceImpl{UserRepo: userRepo}
}
