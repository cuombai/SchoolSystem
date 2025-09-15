package models

import (
    "time"

    "go.mongodb.org/mongo-driver/bson/primitive"
)

type Transcript struct {
    ID        primitive.ObjectID `bson:"_id,omitempty"`
    StudentID primitive.ObjectID `bson:"student_id"`
    Term      string             `bson:"term"`
    FilePath  string             `bson:"file_path"` // If stored as file
    GeneratedAt time.Time        `bson:"generated_at"`
}
