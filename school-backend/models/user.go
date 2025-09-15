package models

import (
    "time"

    "go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
    ID        primitive.ObjectID `bson:"_id,omitempty"`
    Name      string             `bson:"name"`
    Email     string             `bson:"email"`
    Password  string             `bson:"password"`
    Role      string             `bson:"role"` // "student", "teacher", "admin"
    Phone     string             `bson:"phone,omitempty"`
    CreatedAt time.Time          `bson:"created_at"`
}
