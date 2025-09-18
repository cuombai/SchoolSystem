package handler

import (
	"net/http"
    "log"
	"schoolsystem/school-backend/internal/auth"
	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
    var body struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    if err := c.BindJSON(&body); err != nil {
        log.Println("Failed to bind JSON:", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    log.Println("Login attempt for:", body.Email)

    // Query user from DB
    var user User // your user model
    err := collection.FindOne(ctx, bson.M{"email": body.Email}).Decode(&user)
    if err != nil {
        log.Println("User not found:", err)
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
        return
    }

    // Compare passwords
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
        log.Println("Password mismatch:", err)
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
        return
    }

    // Generate token
    token, err := auth.GenerateToken(user.ID.Hex(), user.Role)
    if err != nil {
        log.Println("Token generation failed:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Token error"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}
