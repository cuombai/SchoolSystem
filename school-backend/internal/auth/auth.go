package auth

import (
    "errors"
    // "schoolsystem/school-backend/models"
    "schoolsystem/school-backend/repository"
)

func Authenticate(req LoginRequest) (string, error) {
    user, err := repository.FindUserByEmail(req.Email)
    if err != nil {
        return "", errors.New("user not found")
    }

    if !CheckPasswordHash(req.Password, user.Password) {
        return "", errors.New("invalid password")
    }

    token, err := GenerateToken(user.ID.Hex(), user.Role)
    if err != nil {
        return "", errors.New("token generation failed")
    }

    return token, nil
}
