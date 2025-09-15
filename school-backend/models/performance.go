package models

import (
    "time"

    "go.mongodb.org/mongo-driver/bson/primitive"
)

type Performance struct {
    ID        primitive.ObjectID `bson:"_id,omitempty"`
    StudentID primitive.ObjectID `bson:"student_id"`
    Subject   string             `bson:"subject"`
    Grades    int                `bson:"grades"`
    Remarks   string             `bson:"remarks"`
    Term      string             `bson:"term"`
    CreatedAt time.Time          `bson:"created_at"`
}
