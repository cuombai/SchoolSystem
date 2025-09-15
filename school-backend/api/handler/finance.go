package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetFeeBalance(c *gin.Context) {
	//TODO: Call internal.finance.getFeebalance
	c.JSON(http.StatusOK, gin.H{
		"message": "Fee balance",
	})
}