package db

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func Init(uri string, db *mongo.Client) error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	uri = os.Getenv("MONGO_URI")
	if uri == "" {
		log.Fatal("Mongo URI not set in .env")
	}
	db, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri).SetServerAPIOptions(options.ServerAPI(options.ServerAPIVersion1)))
	if err != nil {
		return err
	}
	return nil
}
