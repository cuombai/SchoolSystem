package handler

import (
    "context"
    "log"
    "net/http"
    "time"

    "schoolsystem/school-backend/internal/auth"
    "go.mongodb.org/mongo-driver/bson/primitive"


    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "golang.org/x/crypto/bcrypt"
)

// User model (adjust fields as needed)
type User struct {
    ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    Email    string             `bson:"email" json:"email"`
    Password string             `bson:"password" json:"password"`
    Role     string             `bson:"role" json:"role"`
}

var (
    collection *mongo.Collection
    ctx        context.Context
)

func init() {
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

    db := client.Database("schoolDB") // Replace with your actual DB name
    collection = db.Collection("users") // Replace with your actual collection name
}

func LoginHandler(c *gin.Context) {
    var body struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    if err := c.BindJSON(&body); err != nil {
        log.Println("Failed to bind JSON:", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    log.Println("Login attempt for:", body.Email)

    var user User
    err := collection.FindOne(ctx, bson.M{"email": body.Email}).Decode(&user)
    if err != nil {
        log.Println("User not found:", err)
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
        return
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
        log.Println("Password mismatch:", err)
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
        return
    }

    token, err := auth.GenerateToken(user.ID.Hex(), user.Role)
    if err != nil {
        log.Println("Token generation failed:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Token error"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}
