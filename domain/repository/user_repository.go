package repository

type UserRepository interface {
	AddFollower(userID, targetID string) error
	GetFollowers(userID string) ([]string, error)
}
