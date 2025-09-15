package repository

import (
    "schoolsystem/school-backend/config"
    "schoolsystem/school-backend/models"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

func GetFeeRecords(studentID primitive.ObjectID) ([]models.FeeRecord, error) {
    ctx, cancel := GetContext()
    defer cancel()

    cursor, err := config.DB.Collection("fees").Find(ctx, bson.M{"student_id": studentID})
    if err != nil {
        return nil, err
    }

    var results []models.FeeRecord
    if err := cursor.All(ctx, &results); err != nil {
        return nil, err
    }
    return results, nil
}

func InsertFeeRecord(record models.FeeRecord) error {
    ctx, cancel := GetContext()
    defer cancel()

    _, err := config.DB.Collection("fees").InsertOne(ctx, record)
    return err
}
