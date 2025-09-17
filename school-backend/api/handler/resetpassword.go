package handler

import (
	"context"
	"net/http"
	"schoolsystem/school-backend/config"
	"schoolsystem/school-backend/internal/auth"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func ResetPasswordHandler(c *gin.Context){
	token := c.Param("token")
	
	var body struct {
		Password string `json:"password"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})

		return

	}


	//Decode token to get userID
	payload, err := auth.DecodeResetToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid or Expired token",
		})
		return
	}

	//Hash new password
	hashed, err := auth.HashPassword(body.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Password hashing failed",
		})
		return
	}

	//Update password in DB
	collection := config.DB.Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = collection.UpdateOne(ctx,
		bson.M{"_id": payload.UserID},
		bson.M{"$set": bson.M{"password": hashed}},
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update password",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Password updated successfully",
	})


}