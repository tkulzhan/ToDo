package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client = mongoClient()

func mongoClient() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	dbUsername := GetEnv("DB_USERNAME", "tkulzhan")
	dbPassword := GetEnv("DB_PASSWORD", "tkulzhan")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("mongodb+srv://%s:%s@cluster.czbqaif.mongodb.net/?retryWrites=true&w=majority", dbUsername, dbPassword)))
	if err != nil {
		log.Println("ERROR: " + err.Error())
	}
	return client
}

func GetEnv(key, fallback string) string {
	err := godotenv.Load("./app.env")
	if err != nil {
		log.Println("Error loading app.env: ", err)
	}
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	log.Println("Could not find " + key + " in env. Returning fallback")
	return fallback
}
