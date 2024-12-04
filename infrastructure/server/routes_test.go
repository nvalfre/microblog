package server

import (
	"bytes"
	"github.com/stretchr/testify/mock"
	"microblog/domain/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) AddFollower(userID, targetID string) error {
	args := m.Called(userID, targetID)
	return args.Error(0)
}

func (m *MockUserRepository) GetFollowers(userID string) ([]string, error) {
	args := m.Called(userID)
	return args.Get(0).([]string), args.Error(1)
}

type MockTweetRepository struct {
	mock.Mock
}

func (m *MockTweetRepository) Save(tweet *models.Tweet) error {
	args := m.Called(tweet)
	return args.Error(0)
}

func (m *MockTweetRepository) GetTimeline(userIDs []string) ([]models.Tweet, error) {
	args := m.Called(userIDs)
	return args.Get(0).([]models.Tweet), args.Error(1)
}

type MockCache struct {
	mock.Mock
}

func (m *MockCache) Get(key string) (string, error) {
	args := m.Called(key)
	return args.String(0), args.Error(1)
}

func (m *MockCache) Set(key, value string, expiration int) error {
	args := m.Called(key, value, expiration)
	return args.Error(0)
}

func (m *MockCache) Delete(key string) error {
	args := m.Called(key)
	return args.Error(0)
}

func TestRegisterRoutes(t *testing.T) {
	tests := []struct {
		name       string
		route      string
		method     string
		body       string
		wantStatus int
	}{
		{
			name:       "Generate Token Route",
			route:      "/generate_token?user_id=1",
			method:     http.MethodGet,
			wantStatus: http.StatusOK,
		},
		{
			name:       "Follow User Route",
			route:      "/api/userCollection/follow",
			method:     http.MethodPost,
			body:       `{"followed_user_id": "user123"}`,
			wantStatus: http.StatusUnauthorized,
		},
		{
			name:       "Publish Tweet Route",
			route:      "/api/tweet/",
			method:     http.MethodPost,
			body:       `{"content": "Hello, world!"}`,
			wantStatus: http.StatusUnauthorized,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := gin.Default()

			mockUserRepo := &MockUserRepository{}
			mockTweetRepo := &MockTweetRepository{}
			mockCache := &MockCache{}

			RegisterRoutes(router, mockUserRepo, mockTweetRepo, mockCache)

			var req *http.Request
			if tt.body != "" {
				req = httptest.NewRequest(tt.method, tt.route, bytes.NewBuffer([]byte(tt.body)))
				req.Header.Set("Content-Type", "application/json")
			} else {
				req = httptest.NewRequest(tt.method, tt.route, nil)
			}

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("RegisterRoutes() for %s = %v, want %v", tt.route, w.Code, tt.wantStatus)
			}
		})
	}
}
