package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func e(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{
		"error": message,
	})
}

func success(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
