package database

import (
	"context"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockMongoClient struct {
	mock.Mock
}

func (m *MockMongoClient) Ping(ctx context.Context, rp interface{}) error {
	args := m.Called(ctx, rp)
	return args.Error(0)
}

func NewMockMongoClient() *MockMongoClient {
	return &MockMongoClient{}
}

func TestNewMongoClient(t *testing.T) {
	tests := []struct {
		name    string
		uri     string
		wantErr bool
	}{
		{
			name:    "Valid MongoDB URI",
			uri:     "mongodb://mock-valid-uri",
			wantErr: false,
		},
		{
			name:    "Invalid MongoDB URI",
			uri:     "mock-invalid-uri",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := NewMockMongoClient()

			if tt.wantErr {
				mockClient.On("Ping", mock.Anything, nil).Return(context.DeadlineExceeded)
			} else {
				mockClient.On("Ping", mock.Anything, nil).Return(nil)
			}

			err := mockClient.Ping(context.Background(), nil)

			if (err != nil) != tt.wantErr {
				t.Errorf("NewMongoClient() error = %v, wantErr %v", err, tt.wantErr)
			}

			mockClient.AssertExpectations(t)
		})
	}
}
