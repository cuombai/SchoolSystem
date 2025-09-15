package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func MarkAttendance(c *gin.Context) {
	//TODO: Call internal.teacher.MarkAttendance
	c.JSON(http.StatusOK, gin.H{
		"message": "Attendance marked",
	})
}

func UploadGrades(c *gin.Context) {
	//TODO: Call internal.teacher.UploadMarks
	c.JSON(http.StatusOK, gin.H{
		"message": "Grades uploaded",
	})

}

func UploadRemarks(c *gin.Context) {
	//TODO: Call internal.teacher.UploadRemarks
	c.JSON(http.StatusOK, gin.H{
		"message": "Remarks uploaded",
	})
}

func GetClassList(c *gin.Context) {
	classID := c.Param("id")
	//TODO: Fetch class list
	c.JSON(http.StatusOK, gin.H{
		"message": "class list for " + classID,
	})
}