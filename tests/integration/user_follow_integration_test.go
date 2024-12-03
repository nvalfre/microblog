package integration

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/go-redis/redismock/v9"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"microblog/infrastructure/adapters/cache"
	"microblog/infrastructure/server"
	"microblog/interface/dto"
	"microblog/security/auth"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFollowUserEndpointWithEmbeddedMongoDB(t *testing.T) {
	ctx := context.Background()

	mongoContainer, err := mongodb.Run(ctx, "mongo:latest",
		testcontainers.WithHostPortAccess(27017),
	)
	require.NoError(t, err)
	defer mongoContainer.Terminate(ctx)

	mongoURI, err := mongoContainer.ConnectionString(ctx)
	require.NoError(t, err)

	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	require.NoError(t, err)
	defer mongoClient.Disconnect(ctx)

	redisClient, redisMock := redismock.NewClientMock()
	mockRedisCache := cache.NewRedisCache(redisClient)

	userTimelineCollection := mongoClient.Database("microblog").Collection("user_timeline")

	router := server.NewHTTPServer(mongoClient, mockRedisCache)
	testServer := httptest.NewServer(router)
	defer testServer.Close()

	validToken, err := auth.GenerateSignedToken("1")
	require.NoError(t, err)

	payload := dto.FollowUserRequest{
		UserID:   "1",
		TargetID: "2",
	}
	payloadJSON, err := json.Marshal(payload)
	require.NoError(t, err)

	req, err := http.NewRequest("POST", testServer.URL+"/api/userCollection/follow", bytes.NewBuffer(payloadJSON))
	require.NoError(t, err)
	req.Header.Set("Authorization", validToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)

	require.Equal(t, http.StatusOK, resp.StatusCode)

	require.NoError(t, redisMock.ExpectationsWereMet())

	var result bson.M
	err = userTimelineCollection.FindOne(ctx, bson.M{"userId": "1", "followerId": "2"}).Decode(&result)
	require.NoError(t, err)
	require.Equal(t, "1", result["userId"])
	require.Equal(t, "2", result["followerId"])
}
