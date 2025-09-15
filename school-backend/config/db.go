package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDB() error {
	uri := GetEnv("MONGO_URI", "")
	dbName := GetEnv("DB_NAME", "schoolDB")

	clientOpts := options.Client().ApplyURI(uri)
	client, err := mongo.NewClient(clientOpts)

	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Connect(ctx); err != nil {
		return err
	}

	DB = client.Database(dbName)
	log.Println("Connected to MongoDB")
	return nil
}