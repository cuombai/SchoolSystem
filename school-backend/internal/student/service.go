package student

func GetTranscript(studentID, studentName string) ([]byte, error) {
    results, err := GetStudentPerformance(studentID)
    if err != nil {
        return nil, err
    }
    return GenerateTranscriptPDF(results, studentName)
}
