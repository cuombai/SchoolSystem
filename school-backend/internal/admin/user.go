package admin

import (
    "time"
    "schoolsystem/school-backend/models"
    "schoolsystem/school-backend/repository"
    "schoolsystem/school-backend/internal/auth"

    "go.mongodb.org/mongo-driver/bson/primitive"
)

func AddUser(req CreateUserRequest, adminID string) error {
    hashed, err := auth.HashPassword(req.Password)
    if err != nil {
        return err
    }

    user := models.User{
        Name:      req.Name,
        Email:     req.Email,
        Password:  hashed,
        Role:      req.Role,
        Phone:     req.Phone,
        CreatedAt: time.Now(),
    }

    if err := repository.InsertUser(user); err != nil {
        return err
    }

    // Log action
    return repository.InsertAuditLog(models.AuditLog{
        Action:    "AddUser",
        ActorID:   adminID,
        Timestamp: time.Now(),
        Details:   "Created user: " + req.Email,
    })
}

func DeleteUser(userID string, adminID string) error {
    objID, err := primitive.ObjectIDFromHex(userID)
    if err != nil {
        return err
    }

    if err := repository.DeleteUserByID(objID); err != nil {
        return err
    }

    // Log action
    return repository.InsertAuditLog(models.AuditLog{
        Action:    "DeleteUser",
        ActorID:   adminID,
        Timestamp: time.Now(),
        Details:   "Deleted user ID: " + userID,
    })
}
