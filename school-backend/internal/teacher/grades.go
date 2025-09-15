package teacher

import (
    "time"
    "schoolsystem/school-backend/models"
    "schoolsystem/school-backend/repository"

    "go.mongodb.org/mongo-driver/bson/primitive"
)

func UploadGrades(req MarksRequest) error {
    studentObjID, err := primitive.ObjectIDFromHex(req.StudentID)
    if err != nil {
        return err
    }

    record := models.Performance{
        StudentID: studentObjID,
        Subject:   req.Subject,
        Grades:     req.Grades,
        Remarks:   req.Remarks,
        Term:      req.Term,
        CreatedAt: time.Now(),
    }

    return repository.InsertPerformance(record)
}
