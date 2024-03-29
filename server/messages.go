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
		failure(c, http.StatusBadRequest, err.Error())
		return
	}
	var (
		member = &db.Member{}
		userID = c.MustGet(contextUser).(*db.User).ID
	)
	if err := s.conn.
		Joins("User").
		Joins("Room").
		Find(member, "user_id = ? AND room_id = ?", userID, params.RoomID).
		Error; err != nil {
		failure(c, http.StatusUnauthorized, err.Error())
		return
	}
	message := &db.Message{
		MemberID:     member.ID,
		Member:       member,
		Body:         params.Body,
		CreationDate: time.Now(),
	}
	if err := s.conn.Save(message).Error; err != nil {
		failure(c, http.StatusInternalServerError, err.Error())
		return
	}
	members := []*db.Member{}
	if err := s.conn.Find(&members, "room_id = ?", member.RoomID).Error; err != nil {
		failure(c, http.StatusInternalServerError, err.Error())
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
		failure(c, http.StatusInternalServerError, err.Error())
		return
	}
	s.herald.Send(m, clients)
	success(c)
}
