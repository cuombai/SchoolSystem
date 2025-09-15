package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetStudentPerformance(c *gin.Context) {
	//TODO: call internal.student.GetPerformance
	c.JSON(http.StatusOK, gin.H{
		"message": "Student performance",
	})
}

func DownloadTranscript(c *gin.Context) {
	//TODO: Generate and Stream PDF
	c.JSON(http.StatusOK, gin.H{
		"message" : "Transcript download",
	})
}

func GetFeeSummary(c *gin.Context) {
	//TODO: call internal.finance.GetFeeSummary
	c.JSON(http.StatusOK, gin.H {
		"message": "Fee summary",
	})
}

func DownloadIncvoice(c *gin.Context) {
	term := c.Param("term")
	//TODO: Generate invoice PDF for the term
	c.JSON(http.StatusOK, gin.H{
		"message": "Invoice for term " + term,
	})
}