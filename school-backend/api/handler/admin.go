package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
)


func AddUser(c *gin.Context) {
	//TODO: Call internal.admin.AddUser
	c.JSON(http.StatusOK, gin.H{
		"message": " User added",
	})
}

func DeleteUser(c *gin.Context) {
	userID := c.Param("id")
	//TODO: Call internal.admin.DeleteUser
	c.JSON(http.StatusOK, gin.H {
		"message": "User deleted: " + userID,
	})
}

func GetAuditLogs(c *gin.Context) {
	//TODO: fetch audit logs
	c.JSON(http.StatusOK, gin.H{
		"message": "Audit Logs",
	})
}