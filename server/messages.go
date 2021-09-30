package server

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nathan-osman/go-herald"
	"github.com/nathan-osman/pratl/db"
)

type messages_POST_params struct {
	RoomID int64  `json:"room_id"`
	Body   string `json:"body"`
}

func (s *Server) messages_POST(c *gin.Context) {
	params := &messages_POST_params{}
	if err := c.ShouldBindJSON(params); err != nil {
		c.JSON(http.StatusBadRequest, e(err))
		return
	}
	member := &db.Member{
		UserID: 0,
		RoomID: params.RoomID,
	}
	if err := s.conn.Joins("User").Where(member).Find(member).Error; err != nil {
		c.JSON(http.StatusUnauthorized, e(err))
		return
	}
	message := &db.Message{
		MemberID:     member.ID,
		Member:       member,
		Body:         params.Body,
		CreationDate: time.Now(),
	}
	if err := s.conn.Save(message).Error; err != nil {
		c.JSON(http.StatusInternalServerError, e(err))
		return
	}
	members := []*db.Member{}
	if err := s.conn.Where("room_id", member.RoomID).Find(&members).Error; err != nil {
		c.JSON(http.StatusInternalServerError, e(err))
		return
	}
	clients := []*herald.Client{}
	for _, m := range members {
		for _, c := range s.herald.Clients() {
			if m.UserID == c.Data.(int64) {
				clients = append(clients, c)
			}
		}
	}
	m, err := herald.NewMessage(messageMessage, message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, e(err))
		return
	}
	s.herald.Send(m, clients)
	c.JSON(http.StatusOK, gin.H{})
}
