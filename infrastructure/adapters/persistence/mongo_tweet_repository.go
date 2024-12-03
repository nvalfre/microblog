package persistence

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"microblog/domain/models"
	"microblog/infrastructure/logger"
)

type MongoTweetRepository struct {
	Collection *mongo.Collection
}

func (r *MongoTweetRepository) Save(tweet *models.Tweet) error {
	_, err := r.Collection.InsertOne(context.Background(), tweet)
	return err
}

func (r *MongoTweetRepository) GetTimeline(userIDs []string) ([]models.Tweet, error) {
	filter := bson.M{"user_id": bson.M{"$in": userIDs}}
	opts := options.Find().SetSort(bson.D{{"created_at", -1}}).SetLimit(50)

	cursor, err := r.Collection.Find(context.Background(), filter, opts)
	if err != nil {
		logger.Error("Error retrieving timeline", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	var tweets []models.Tweet
	if err := cursor.All(context.Background(), &tweets); err != nil {
		logger.Error("Error decoding tweets", err)
		return nil, err
	}
	return tweets, nil
}

// NewMongoTweetRepository creates a new instance of MongoTweetRepository.
func NewMongoTweetRepository(collection *mongo.Collection) *MongoTweetRepository {
	return &MongoTweetRepository{
		Collection: collection,
	}
}
