package tweet

import (
	"errors"
	"microblog/domain/models"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// MockTweetService is a mock implementation of TweetService.
type MockTweetPublishService struct {
	mock.Mock
}

func (m *MockTweetPublishService) SaveTweet(tweet *models.Tweet) error {
	args := m.Called(tweet)
	return args.Error(0)
}

func (m *MockTweetPublishService) GetTimeline(userIDs []string) ([]models.Tweet, error) {
	return nil, nil
}

func TestPublishTweetUseCase_Execute(t *testing.T) {
	mockTweetService := new(MockTweetPublishService)

	useCase := &PublishTweetUseCase{
		TweetService: mockTweetService,
	}

	tests := []struct {
		name    string
		tweet   *models.Tweet
		mockErr error
		wantErr bool
	}{
		{
			name: "success - tweet saved",
			tweet: &models.Tweet{
				ID:      "tweet1",
				UserID:  "user1",
				Content: "Hello World",
			},
			mockErr: nil,
			wantErr: false,
		},
		{
			name: "failure - service error",
			tweet: &models.Tweet{
				ID:      "tweet2",
				UserID:  "user2",
				Content: "Another Tweet",
			},
			mockErr: errors.New("failed to save tweet"),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockTweetService.On("SaveTweet", tt.tweet).Return(tt.mockErr).Once()

			err := useCase.Execute(tt.tweet)

			if tt.wantErr {
				require.Error(t, err)
				require.Equal(t, tt.mockErr, err)
			} else {
				require.NoError(t, err)
			}

			mockTweetService.AssertExpectations(t)
		})
	}
}
