package server

import (
	"github.com/gin-gonic/gin"
)

func e(err error) gin.H {
	return gin.H{"error": err.Error()}
}
