package server

import (
	"github.com/gin-gonic/gin"
	"github.com/nathan-osman/go-herald"
	"github.com/nathan-osman/pratl/db"
)

const (

	// messageRooms contains a list of rooms the user is in
	messageRooms = "rooms"

	// messageMessage indicates that a new message has been sent
	messageMessage = "message"

	// messageError indicates an error condition
	messageError = "error"
)

func (s *Server) ws(c *gin.Context) {
	client, err := s.herald.AddClient(
		c.Writer,
		c.Request,
		c.MustGet(contextUser).(*db.User).ID,
	)
	if err != nil {
		s.logger.Error().Msg(err.Error())
		return
	}
	client.Wait()
}

func (s *Server) processMessage(m *herald.Message, c *herald.Client) {
	// TODO: handle incoming messages
}

func (s *Server) processClientAdded(c *herald.Client) {
	var (
		userID  = c.Data.(int64)
		members = []*db.Member{}
	)
	if err := s.conn.Joins("Room").
		Limit(10).
		Find(&members, "user_id = ?", userID).
		Error; err != nil {
		s.sendErrorMessage(err.Error(), c)
	}
	s.herald.Send(
		s.mustNewMessage(messageRooms, members),
		[]*herald.Client{c},
	)
}

func (s *Server) processClientRemoved(c *herald.Client) {
	// TODO: handle user disconnecting
}
