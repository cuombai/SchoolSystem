package finance

import (
    // "schoolsystem/school-backend/models"
    "schoolsystem/school-backend/repository"

    "go.mongodb.org/mongo-driver/bson/primitive"
)

func GetFeeBalance(studentID string) ([]FeeInvoice, error) {
    objID, err := primitive.ObjectIDFromHex(studentID)
    if err != nil {
        return nil, err
    }

    records, err := repository.GetFeeRecords(objID)
    if err != nil {
        return nil, err
    }

    var invoices []FeeInvoice
    for _, r := range records {
        invoices = append(invoices, FeeInvoice{
            Term:        r.Term,
            InvoiceNo:   r.InvoiceNo,
            AmountDue:   r.AmountDue,
            AmountPaid:  r.AmountPaid,
            Balance:     r.AmountDue - r.AmountPaid,
            DueDate:     r.DueDate.Format("2006-01-02"),
        })
    }
    return invoices, nil
}
