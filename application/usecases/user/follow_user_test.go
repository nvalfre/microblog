package user

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type MockUserFollowService struct {
	mock.Mock
}

func (m *MockUserFollowService) AddFollower(userID, followerID string) error {
	args := m.Called(userID, followerID)
	return args.Error(0)
}

func (m *MockUserFollowService) GetFollowers(userID string) ([]string, error) {
	args := m.Called(userID)
	return args.Get(0).([]string), args.Error(1)
}

func TestFollowUserUseCase_Execute(t *testing.T) {
	mockRepo := new(MockUserFollowService)

	useCase := &FollowUserUseCase{
		UserService: mockRepo,
	}

	tests := []struct {
		name     string
		userID   string
		targetID string
		mockErr  error
		wantErr  bool
	}{
		{
			name:     "success - follow user",
			userID:   "user1",
			targetID: "user2",
			mockErr:  nil,
			wantErr:  false,
		},
		{
			name:     "failure - repository error",
			userID:   "user3",
			targetID: "user4",
			mockErr:  errors.New("failed to add follower"),
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.On("AddFollower", tt.userID, tt.targetID).Return(tt.mockErr).Once()

			err := useCase.Execute(tt.userID, tt.targetID)

			if tt.wantErr {
				require.Error(t, err)
				require.Equal(t, tt.mockErr, err)
			} else {
				require.NoError(t, err)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}
