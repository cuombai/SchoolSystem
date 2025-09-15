package finance

type FeeInvoice struct {
    StudentName string  `json:"student_name"`
    Term        string  `json:"term"`
    InvoiceNo   string  `json:"invoice_no"`
    AmountDue   float64 `json:"amount_due"`
    AmountPaid  float64 `json:"amount_paid"`
    Balance     float64 `json:"balance"`
    DueDate     string  `json:"due_date"`
}
