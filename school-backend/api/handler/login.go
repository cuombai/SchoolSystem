package handler

import (
	"net/http"
	"schoolsystem/school-backend/internal/auth"
	"schoolsystem/school-backend/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"log"

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

    var user models.User
    err := collection.FindOne(ctx, bson.M{"email": body.Email}).Decode(&user)
    if err != nil {
        log.Println("User not found:", err)
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
        return
    }

	if user.Password != body.Password{
		log.Println("Password mismatch")
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Authentication failed",
		})
	}

    // if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
    //     log.Println("Password mismatch:", err)
    //     c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
    //     return
    // }

    token, err := auth.GenerateToken(user.ID.Hex(), user.Role)
    if err != nil {
        log.Println("Token generation failed:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Token error"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}
