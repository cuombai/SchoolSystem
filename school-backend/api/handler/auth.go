package handler

import (
	"context"
	"log"
	// "schoolsystem/school-backend/config"

	// "net/http"
	"time"

	// "schoolsystem/school-backend/internal/auth"
	// "go.mongodb.org/mongo-driver/bson/primitive"

	// "github.com/gin-gonic/gin"
	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "golang.org/x/crypto/bcrypt"
)

// User model (adjust fields as needed)
// type User struct {
//     ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
//     Email    string             `bson:"email" json:"email"`
//     Password string             `bson:"password" json:"password"`
//     Role     string             `bson:"role" json:"role"`
// }

var (
    collection *mongo.Collection 
    ctx        context.Context 
)

func Init(db *mongo.Database) {
    client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
    if err != nil {
        log.Fatal("Mongo client creation failed:", err)
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    err = client.Connect(ctx)
    if err != nil {
        log.Fatal("Mongo connection failed:", err)
    }

    // DB := client.Database("school") // Replace with your actual DB name
    collection = db.Collection("users") // Replace with your actual collection name
}
