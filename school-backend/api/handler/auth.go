package handler

import (
	"net/http"
	"schoolsystem/school-backend/internal/auth"
	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	var credentials auth.LoginRequest
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	token, err := auth.Authenticate(credentials)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Authentication failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}