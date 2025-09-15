package repository

import (
    "schoolsystem/school-backend/config"
    "schoolsystem/school-backend/models"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

func FindUserByEmail(email string) (*models.User, error) {
    ctx, cancel := GetContext()
    defer cancel()

    var user models.User
    err := config.DB.Collection("users").FindOne(ctx, bson.M{"email": email}).Decode(&user)
    return &user, err
}

func InsertUser(user models.User) error {
    ctx, cancel := GetContext()
    defer cancel()

    _, err := config.DB.Collection("users").InsertOne(ctx, user)
    return err
}

func DeleteUserByID(id primitive.ObjectID) error {
    ctx, cancel := GetContext()
    defer cancel()

    _, err := config.DB.Collection("users").DeleteOne(ctx, bson.M{"_id": id})
    return err
}
