package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nathan-osman/go-herald"
)

func (s *Server) mustNewMessage(messageType string, data interface{}) *herald.Message {
	m, err := herald.NewMessage(messageType, data)
	if err != nil {
		s.logger.Error().Msgf("NewMessage() failed: %s", err.Error())
	}
	return m
}

func (s *Server) sendErrorMessage(message string, client *herald.Client) {
	s.herald.Send(
		s.mustNewMessage(messageError, message),
		[]*herald.Client{client},
	)
}

func failure(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{
		"error": message,
	})
}

func success(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
