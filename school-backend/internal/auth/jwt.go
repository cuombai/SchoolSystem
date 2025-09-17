package auth

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func GenerateToken(userID string, role string) (string, error) {
    claims := jwt.MapClaims{
        "userID": userID,
        "role":   role,
        "exp":    time.Now().Add(24 * time.Hour).Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtSecret)
}

type ResetPayload struct {
    UserID string `json:"userID"`
    jwt.RegisteredClaims
}

func DecodeResetToken(tokenStr string)(*ResetPayload, error) {
    token, err := jwt.ParseWithClaims(tokenStr, &ResetPayload{}, func(token *jwt.Token)(interface{}, error){
        return []byte(os.Getenv("JWT_SECRET")), nil
    })

    if err != nil || !token.Valid {
        return nil, err
    }

    payload, ok := token.Claims.(*ResetPayload)

    if !ok {
        return nil, errors.New("Invalid token payload")
    }

    return payload, nil
}