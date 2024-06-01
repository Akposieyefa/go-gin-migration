package utils

import (
	"github.com/gin-gonic/gin"
)

func WriteSuccess(c *gin.Context, status int, message string, data any, success bool) {
	c.JSON(status, gin.H{
		"message": message,
		"data":    data,
		"success": success,
	})
}

func WriteError(c *gin.Context, status int, message string, success bool) {
	c.JSON(status, gin.H{
		"message": message,
		"success": success,
	})
}
