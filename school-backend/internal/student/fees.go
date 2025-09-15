package student

import (
    // "schoolsystem/school-backend/models"
    "schoolsystem/school-backend/repository"

    "go.mongodb.org/mongo-driver/bson/primitive"
)

func GetFeeSummary(studentID string) ([]FeeSummary, error) {
    objID, err := primitive.ObjectIDFromHex(studentID)
    if err != nil {
        return nil, err
    }

    records, err := repository.GetFeeRecords(objID)
    if err != nil {
        return nil, err
    }

    var summary []FeeSummary
    for _, r := range records {
        summary = append(summary, FeeSummary{
            Term:       r.Term,
            InvoiceNo:  r.InvoiceNo,
            AmountDue:  r.AmountDue,
            AmountPaid: r.AmountPaid,
            Balance:    r.AmountDue - r.AmountPaid,
            DueDate:    r.DueDate.Format("2006-01-02"),
        })
    }
    return summary, nil
}
