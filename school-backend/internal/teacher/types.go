package teacher

type AttendanceRequest struct {
    StudentID string `json:"student_id"`
    Date      string `json:"date"`      // ISO format
    Status    string `json:"status"`    // "present" or "absent"
}

type MarksRequest struct {
    StudentID string `json:"student_id"`
    Subject   string `json:"subject"`
    Grades     int    `json:"grades"`
    Remarks   string `json:"remarks"`
    Term      string `json:"term"`
}

type ClassStudent struct {
    ID    string `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}
