package server

import (
	"github.com/nathan-osman/go-herald"
)

const (
	// messageMessage indicates that a new message has been sent
	messageMessage = "message"
)

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
