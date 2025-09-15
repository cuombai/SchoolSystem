package teacher

import (
    "schoolsystem/school-backend/models"
    "schoolsystem/school-backend/config"

    "go.mongodb.org/mongo-driver/bson"
)

func GetClassList(classID string) ([]ClassStudent, error) {
    ctx, cancel := config.GetContext()
    defer cancel()

    cursor, err := config.DB.Collection("users").Find(ctx, bson.M{
        "role":     "student",
        "class_id": classID,
    })
    if err != nil {
        return nil, err
    }

    var students []models.User
    if err := cursor.All(ctx, &students); err != nil {
        return nil, err
    }

    var result []ClassStudent
    for _, s := range students {
        result = append(result, ClassStudent{
            ID:    s.ID.Hex(),
            Name:  s.Name,
            Email: s.Email,
        })
    }
    return result, nil
}
