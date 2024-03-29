package server

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nathan-osman/pratl/db"
)

type rooms_POST_params struct {
	Name string `json:"name"`
}

func (s *Server) rooms_POST(c *gin.Context) {
	params := &rooms_POST_params{}
	if err := c.ShouldBindJSON(params); err != nil {
		failure(c, http.StatusBadRequest, err.Error())
		return
	}
	room := &db.Room{
		Name:         params.Name,
		CreationDate: time.Now(),
	}
	if err := s.conn.Save(room).Error; err != nil {
		failure(c, http.StatusInternalServerError, err.Error())
		return
	}
	var (
		userID = c.MustGet(contextUser).(*db.User).ID
		member = &db.Member{
			UserID:  userID,
			RoomID:  room.ID,
			IsOwner: true,
			IsAdmin: true,
		}
	)
	if err := s.conn.Save(member).Error; err != nil {
		failure(c, http.StatusInternalServerError, err.Error())
		return
	}
	success(c)
}
