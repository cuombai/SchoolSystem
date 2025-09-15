package repository

import (
	"schoolsystem/school-backend/config"
    "schoolsystem/school-backend/models"

    "go.mongodb.org/mongo-driver/bson"
)

func InsertAttendance(record models.Attendance) error {
    ctx, cancel := GetContext()
    defer cancel()

    _, err := config.DB.Collection("attendance").InsertOne(ctx, record)
    return err
}

func GetAttendanceByStudent(studentID string) ([]models.Attendance, error) {
    ctx, cancel := GetContext()
    defer cancel()

    cursor, err := config.DB.Collection("attendance").Find(ctx, bson.M{"student_id": studentID})
    if err != nil {
        return nil, err
    }

    var results []models.Attendance
    if err := cursor.All(ctx, &results); err != nil {
        return nil, err
    }
    return results, nil
}
