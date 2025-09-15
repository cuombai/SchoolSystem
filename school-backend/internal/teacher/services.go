package teacher

func SubmitMarksAndAttendance(m MarksRequest, a AttendanceRequest, teacherID string) error {
    if err := UploadGrades(m); err != nil {
        return err
    }
    return MarkAttendance(a, teacherID)
}
