package persistence

import (
	"context"
	"github.com/sirupsen/logrus"
	"microblog/infrastructure/logger"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoUserTimelineRepository struct {
	Collection *mongo.Collection
}

func (r *MongoUserTimelineRepository) AddFollower(userID, followerID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"userId": userID, "followerId": followerID}
	update := bson.M{
		"$setOnInsert": bson.M{
			"createdAt": time.Now(),
		},
	}

	_, err := r.Collection.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
	if err != nil {
		logger.Error("Error adding follower to user", err, logrus.Fields{"followerID": followerID, "userID": userID})
		return err
	}

	return nil
}

func (r *MongoUserTimelineRepository) GetFollowers(userID string) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"userId": userID}
	cursor, err := r.Collection.Find(ctx, filter)
	if err != nil {
		logger.Error("Error retrieving followers", err, logrus.Fields{"userID": userID})
		return nil, err
	}

	var followers []string
	for cursor.Next(ctx) {
		var rel struct {
			FollowerId string `bson:"followerId"`
		}
		if err := cursor.Decode(&rel); err == nil {
			followers = append(followers, rel.FollowerId)
		}
	}
	cursor.Close(ctx)
	return followers, nil
}

func NewMongoUserTimelineRepository(collection *mongo.Collection) *MongoUserTimelineRepository {
	return &MongoUserTimelineRepository{
		Collection: collection,
	}
}
