package teacher

import (
    "time"
    "schoolsystem/school-backend/models"
    "schoolsystem/school-backend/repository"

    "go.mongodb.org/mongo-driver/bson/primitive"
)

func MarkAttendance(req AttendanceRequest, teacherID string) error {
    studentObjID, err := primitive.ObjectIDFromHex(req.StudentID)
    if err != nil {
        return err
    }

    teacherObjID, er := primitive.ObjectIDFromHex(teacherID)
    if er != nil {
        return er
    }

    date, err := time.Parse("2006-01-02", req.Date)
    if err != nil {
        return err
    }

    record := models.Attendance{
        StudentID: studentObjID,
        Date:      date,
        Status:    req.Status,
        MarkedBy: teacherObjID,
    }

    return repository.InsertAttendance(record)
}
