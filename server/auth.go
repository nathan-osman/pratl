package server

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/nathan-osman/pratl/db"
)

const (
	sessionUserID = "user_id"
	contextUser   = "user"
)

type auth_login_POST_params struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (s *Server) auth_login_POST(c *gin.Context) {
	params := &auth_login_POST_params{}
	if err := c.ShouldBindJSON(params); err != nil {
		failure(c, http.StatusBadRequest, err.Error())
		return
	}
	user := &db.User{}
	if err := s.conn.First(user, "username = ?", params.Username).Error; err != nil {
		failure(c, http.StatusUnauthorized, err.Error())
		return
	}
	if err := user.Authenticate(params.Password); err != nil {
		failure(c, http.StatusUnauthorized, err.Error())
		return
	}
	session := sessions.Default(c)
	session.Set(sessionUserID, user.ID)
	if err := session.Save(); err != nil {
		failure(c, http.StatusInternalServerError, err.Error())
		return
	}
	success(c)
}

type auth_register_POST_params struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"email"`
}

func (s *Server) auth_register_POST(c *gin.Context) {
	params := &auth_register_POST_params{}
	if err := c.ShouldBindJSON(params); err != nil {
		failure(c, http.StatusBadRequest, err.Error())
		return
	}
	user := &db.User{
		Username: params.Username,
		Email:    params.Email,
		IsActive: true,
	}
	if err := user.SetPassword(params.Password); err != nil {
		failure(c, http.StatusInternalServerError, err.Error())
		return
	}
	if err := s.conn.Save(user).Error; err != nil {
		failure(c, http.StatusInternalServerError, err.Error())
		return
	}
	success(c)
}

func (s *Server) requireLogin(c *gin.Context) {
	session := sessions.Default(c)
	var (
		v    = session.Get(sessionUserID)
		user = &db.User{}
	)
	if err := s.conn.First(user, v).Error; err != nil {
		failure(c, http.StatusUnauthorized, err.Error())
		c.Abort()
		return
	}
	c.Set(contextUser, user)
}
