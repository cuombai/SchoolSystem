package models

import (
    "time"

    "go.mongodb.org/mongo-driver/bson/primitive"
)

type Attendance struct {
    ID        primitive.ObjectID `bson:"_id,omitempty"`
    StudentID primitive.ObjectID `bson:"student_id"`
    Date      time.Time          `bson:"date"`
    Status    string             `bson:"status"` // "present", "absent"
    MarkedBy  primitive.ObjectID `bson:"marked_by"` // Teacher ID
}
