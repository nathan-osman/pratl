package server

import (
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/nathan-osman/pratl/db"
)

type authenticatorParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// authenticator creates a JWT from the provided login credentials.
func (s *Server) authenticator(c *gin.Context) (interface{}, error) {
	params := &authenticatorParams{}
	if err := c.ShouldBindJSON(params); err != nil {
		return nil, err
	}
	u := &db.User{}
	if err := s.conn.First(u, "username = ?", params.Username).Error; err != nil {
		return nil, jwt.ErrFailedAuthentication
	}
	if err := u.Authenticate(params.Password); err != nil {
		return nil, jwt.ErrFailedAuthentication
	}
	return u.ID, nil
}

// authorizator examines the output of identityHandler.
func (s *Server) authorizator(data interface{}, c *gin.Context) bool {
	return data != nil
}

// payloadFunc adds the identity (user ID) to the token
func (s *Server) payloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(int64); ok {
		return jwt.MapClaims{
			identityKey: v,
		}
	}
	return jwt.MapClaims{}
}

// identityHandler loads the User from the database.
func (s *Server) identityHandler(c *gin.Context) interface{} {
	u := &db.User{}
	if err := s.conn.
		First(u, "id = ? AND is_active = ?", jwt.ExtractClaims(c)[identityKey], true).
		Error; err != nil {
		return nil
	}
	return u
}

type auth_register_POST_params struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"email"`
}

func (s *Server) auth_register_POST(c *gin.Context) {
	params := &auth_register_POST_params{}
	if err := c.ShouldBindJSON(params); err != nil {
		e(c, http.StatusBadRequest, err.Error())
		return
	}
	user := &db.User{
		Username: params.Username,
		Email:    params.Email,
		IsActive: true,
	}
	if err := user.SetPassword(params.Password); err != nil {
		e(c, http.StatusInternalServerError, err.Error())
		return
	}
	if err := s.conn.Save(user).Error; err != nil {
		e(c, http.StatusInternalServerError, err.Error())
		return
	}
	success(c)
}
