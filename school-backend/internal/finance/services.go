package finance

import "fmt"

func GetInvoiceForStudent(studentID, studentName, term string) ([]byte, error) {
    invoices, err := GetFeeBalance(studentID)
    if err != nil {
        return nil, err
    }

    var selected *FeeInvoice
    for _, inv := range invoices {
        if inv.Term == term {
            selected = &inv
            break
        }
    }

    if selected == nil {
        return nil, fmt.Errorf("invoice not found for term: %s", term)
    }

    selected.StudentName = studentName
    return GenerateInvoicePDF(*selected)
}
