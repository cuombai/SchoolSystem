package student

import (
    // "schoolsystem/school-backend/models"
    "schoolsystem/school-backend/repository"

    "go.mongodb.org/mongo-driver/bson/primitive"
)

func GetStudentPerformance(studentID string) ([]SubjectResult, error) {
    objID, err := primitive.ObjectIDFromHex(studentID)
    if err != nil {
        return nil, err
    }

    records, err := repository.GetPerformanceByStudent(objID)
    if err != nil {
        return nil, err
    }

    var results []SubjectResult
    for _, r := range records {
        results = append(results, SubjectResult{
            Subject: r.Subject,
            Grades:  r.Grades,
            Remarks: r.Remarks,
            Term:    r.Term,
        })
    }
    return results, nil
}
