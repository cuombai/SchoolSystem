package repository

import (
    "schoolsystem/school-backend/config"
    "schoolsystem/school-backend/models"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

func GetPerformanceByStudent(studentID primitive.ObjectID) ([]models.Performance, error) {
    ctx, cancel := GetContext()
    defer cancel()

    cursor, err := config.DB.Collection("performance").Find(ctx, bson.M{"student_id": studentID})
    if err != nil {
        return nil, err
    }

    var results []models.Performance
    if err := cursor.All(ctx, &results); err != nil {
        return nil, err
    }
    return results, nil
}

func InsertPerformance(record models.Performance) error {
    ctx, cancel := GetContext()
    defer cancel()

    _, err := config.DB.Collection("performance").InsertOne(ctx, record)
    return err
}
