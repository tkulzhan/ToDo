package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client = mongoClient()

func mongoClient() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://tkulzhan:tkulzhan@cluster.czbqaif.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		panic(err)
	}
	return client
}
