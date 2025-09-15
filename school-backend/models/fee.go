package models

import (
    "time"

    "go.mongodb.org/mongo-driver/bson/primitive"
)

type FeeRecord struct {
    ID         primitive.ObjectID `bson:"_id,omitempty"`
    StudentID  primitive.ObjectID `bson:"student_id"`
    Term       string             `bson:"term"`
    InvoiceNo  string             `bson:"invoice_no"`
    AmountDue  float64            `bson:"amount_due"`
    AmountPaid float64            `bson:"amount_paid"`
    DueDate    time.Time          `bson:"due_date"`
    CreatedAt  time.Time          `bson:"created_at"`
}
