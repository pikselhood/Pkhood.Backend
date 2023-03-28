package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"pkhood-backend/src/settings"
	"time"
)

type MongoClient[T any] struct {
	collection *mongo.Collection
}

func NewMongoClient[T any](databaseName string, collectionName string) (mongoClient *MongoClient[T]) {

	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)

	clientOptions := options.Client().
		ApplyURI(settings.MongoConnection).
		SetServerAPIOptions(serverAPIOptions)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database(databaseName).Collection(collectionName)

	return &MongoClient[T]{
		collection: collection,
	}
}

func (mongoClient *MongoClient[T]) Get(filter bson.D) (*T, error) {

	var result T
	if err := mongoClient.collection.FindOne(context.TODO(), filter).Decode(&result); err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			// TODO not found
			return nil, &echo.HTTPError{
				Code:     404,
				Message:  "User not found",
				Internal: nil,
			}
		}

		return nil, err
	}

	return &result, nil
}

func (mongoClient *MongoClient[T]) GetById(id uuid.UUID) (*T, error) {

	filter := bson.D{{"_id", &id}}

	var result T
	if err := mongoClient.collection.FindOne(context.TODO(), filter).Decode(&result); err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return nil, nil
		}

		return nil, err
	}

	return &result, nil
}

func (mongoClient *MongoClient[T]) Insert(document *T) (*mongo.InsertOneResult, error) {

	var (
		result *mongo.InsertOneResult
		err    error
	)

	if result, err = mongoClient.collection.InsertOne(context.TODO(), document); err != nil {
		return nil, err
	}

	return result, nil
}
