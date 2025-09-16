package repository

import (
    "context"
    "log"
    "sync"
    "time"

    "schoolsystem/school-backend/config"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var (
    clientInstance *mongo.Client
    clientOnce     sync.Once
)

// getClient initializes and returns a singleton MongoDB client
func getClient() *mongo.Client {
    clientOnce.Do(func() {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        defer cancel()

        cfg := config.LoadConfig()
        client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoURI))
        if err != nil {
            log.Fatalf("MongoDB connection failed: %v", err)
        }
        clientInstance = client
    })
    return clientInstance
}

// getDB returns the connected MongoDB database
func getDB() *mongo.Database {
    cfg := config.LoadConfig()
    return getClient().Database(cfg.DBName)
}
