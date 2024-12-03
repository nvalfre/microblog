package tweet

import (
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"microblog/domain/models"
	"testing"
)

type MockTweetService struct {
	mock.Mock
}

func (m *MockTweetService) GetTimeline(userIDs []string) ([]models.Tweet, error) {
	args := m.Called(userIDs)
	return args.Get(0).([]models.Tweet), args.Error(1)
}

func (m *MockTweetService) SaveTweet(tweet *models.Tweet) error {
	return nil
}

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) GetFollowers(userID string) ([]string, error) {
	args := m.Called(userID)
	return args.Get(0).([]string), args.Error(1)
}

func (m *MockUserService) AddFollower(userID, followerID string) error {
	return nil
}

type MockCache struct {
	mock.Mock
}

// Get retrieves a value from the cache.
func (c *MockCache) Get(key string) (string, error) {
	args := c.Called(key)
	return args.String(0), args.Error(1)
}

// Set stores a value in the cache with an expiration.
func (c *MockCache) Set(key string, value string, expiration int) error {
	args := c.Called(key, value, expiration)
	return args.Error(0)
}

// Delete removes a value from the cache.
func (c *MockCache) Delete(key string) error {
	return nil
}

func TestGetTimelineUseCase_Execute(t *testing.T) {
	mockTweetService := new(MockTweetService)
	mockUserService := new(MockUserService)
	mockCache := new(MockCache)

	useCase := &GetTimelineUseCase{
		TweetService: mockTweetService,
		UserServuce:  mockUserService,
		Cache:        mockCache,
	}

	tests := []struct {
		name                string
		userID              string
		mockFollowers       []string
		mockTweets          []models.Tweet
		cacheHit            bool
		cachedTimeline      []models.Tweet
		mockCacheError      error
		mockUserServiceErr  error
		mockTweetServiceErr error
		wantErr             bool
	}{
		{
			name:           "success - cache hit",
			userID:         "user1",
			cacheHit:       true,
			cachedTimeline: []models.Tweet{{ID: "tweet1", UserID: "user2", Content: "Cached Tweet"}},
			wantErr:        false,
		},
		{
			name:          "success - cache miss and valid data",
			userID:        "user1",
			cacheHit:      false,
			mockFollowers: []string{"user2", "user3"},
			mockTweets:    []models.Tweet{{ID: "tweet2", UserID: "user2", Content: "New Tweet"}},
			wantErr:       false,
		},
		{
			name:               "failure - user service error",
			userID:             "user1",
			cacheHit:           false,
			mockUserServiceErr: errors.New("user service error"),
			wantErr:            true,
		},
		{
			name:                "failure - tweet service error",
			userID:              "user1",
			cacheHit:            false,
			mockFollowers:       []string{"user2", "user3"},
			mockTweetServiceErr: errors.New("tweet service error"),
			wantErr:             true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cacheKey := "timeline:" + tt.userID

			if tt.cacheHit {
				cachedData, _ := json.Marshal(tt.cachedTimeline)
				mockCache.On("Get", cacheKey).Return(string(cachedData), nil).Once()
			} else {
				mockCache.On("Get", cacheKey).Return("", tt.mockCacheError).Once()

				mockUserService.On("GetFollowers", tt.userID).Return(tt.mockFollowers, tt.mockUserServiceErr).Once()

				if tt.mockUserServiceErr == nil {
					mockTweetService.On("GetTimeline", tt.mockFollowers).Return(tt.mockTweets, tt.mockTweetServiceErr).Once()
				}

				if tt.mockUserServiceErr == nil && tt.mockTweetServiceErr == nil {
					fetchedData, _ := json.Marshal(tt.mockTweets)
					mockCache.On("Set", cacheKey, string(fetchedData), mock.Anything).Return(nil).Once()
				}
			}

			got, err := useCase.Execute(tt.userID)

			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				if tt.cacheHit {
					require.Equal(t, tt.cachedTimeline, got)
				} else {
					require.Equal(t, tt.mockTweets, got)
				}
			}

			mockCache.AssertExpectations(t)
			mockUserService.AssertExpectations(t)
			mockTweetService.AssertExpectations(t)
		})
	}
}
