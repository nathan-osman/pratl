package server

import (
	"github.com/gin-gonic/gin"
	"github.com/nathan-osman/go-herald"
	"github.com/nathan-osman/pratl/db"
)

const (
	// messageMessage indicates that a new message has been sent
	messageMessage = "message"
)

func (s *Server) ws(c *gin.Context) {
	s.herald.AddClient(c.Writer, c.Request, c.MustGet(identityKey).(*db.User).ID)
}

func (s *Server) processMessage(m *herald.Message, c *herald.Client) {
	switch m.Type {
	case messageMessage:
		//...
	}

	// TODO: parse message and store in the database
	// Send to all other clients in the room
}

func (s *Server) processClientAdded(c *herald.Client) {
	//...
}

func (s *Server) processClientRemoved(c *herald.Client) {
	//...
}
