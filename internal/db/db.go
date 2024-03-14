package db

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var db *mongo.Database

func Init() error {
	err := godotenv.Load()
	if err != nil {
		return errors.New("Error loading .env file")
	}

	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		return errors.New("MONGO_URI not set in .env")
	}

	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri).SetServerAPIOptions(options.ServerAPI(options.ServerAPIVersion1)))
	if err != nil {
		return err
	}

	db = client.Database("Tasknight")

	fmt.Println("Connected to db")
	return nil
}

func Close() error {
	return db.Client().Disconnect(context.Background())
}

func GetCollection(coll string) *mongo.Collection {
	return db.Collection(coll)
}
