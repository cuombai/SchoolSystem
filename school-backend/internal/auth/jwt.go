package auth

import (
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
