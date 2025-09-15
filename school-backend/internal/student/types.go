package student

type SubjectResult struct {
    Subject string `json:"subject"`
    Grades  int    `json:"grades"`
    Remarks string `json:"remarks"`
    Term    string `json:"term"`
}

type FeeSummary struct {
    Term        string  `json:"term"`
    InvoiceNo   string  `json:"invoice_no"`
    AmountDue   float64 `json:"amount_due"`
    AmountPaid  float64 `json:"amount_paid"`
    Balance     float64 `json:"balance"`
    DueDate     string  `json:"due_date"`
}
