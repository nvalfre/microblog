package user

type FollowUserPort interface {
	Execute(userID, targetID string) error
}
