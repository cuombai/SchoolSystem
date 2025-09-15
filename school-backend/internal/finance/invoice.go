package finance

import (
	"fmt"
	"schoolsystem/school-backend/utils"
)

func GenerateInvoicePDF(invoice FeeInvoice) ([]byte, error) {
	pdf := utils.NewPDF("Fee Invoice: " + invoice.StudentName)

	pdf.AddLine(fmt.Sprintf("Term: %s", invoice.Term))
	pdf.AddLine(fmt.Sprintf("Invoice No: %s", invoice.InvoiceNo))
	pdf.AddLine(fmt.Sprintf("Amount Due: %.2f", invoice.AmountDue))
	pdf.AddLine(fmt.Sprintf("Amount Paid: %.2f", invoice.AmountPaid))
	pdf.AddLine(fmt.Sprintf("Balance: %.2f", invoice.Balance))
	pdf.AddLine(fmt.Sprintf("Due Date: %s", invoice.DueDate))

	return pdf.Output()
}
