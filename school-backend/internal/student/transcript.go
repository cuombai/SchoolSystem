package student

import (
    // "bytes"
    "fmt"
    "schoolsystem/school-backend/utils"
)

func GenerateTranscriptPDF(results []SubjectResult, studentName string) ([]byte, error) {
    pdf := utils.NewPDF("Transcript for " + studentName)																									
	// pdf.AddLine("Math | 85 marks | Excellent | Term 1")
	// pdf.AddLine("English | 78 marks | Good | Term 1")
	
// Stream or save the PDF


    for _, r := range results {
        line := fmt.Sprintf("%s | %d grades | %s | Term: %s", r.Subject, r.Grades, r.Remarks, r.Term)
        pdf.AddLine(line)
    }

    return pdf.Output()
}
