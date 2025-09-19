package handler

import (
	"net/http"
	"schoolsystem/school-backend/models"

	"github.com/gin-gonic/gin"
)

func SignupHandler(c *gin.Context) {
	var body struct {
		FullName string `json:"fullname"`
		Email string `json:"email"`
		Password string `json: "password"`
		Role string `json: "role"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}

	if body.Role != "student" && body.Role != "teacher" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid role",
		})
		return
	}

	//Hash Password later
	user := models.User {
		Email: body.Email,
		Password: body.Password,
		Role: body.Role,
	}

	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":"Failed to add user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User Created",
	})
}