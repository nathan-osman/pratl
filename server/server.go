package server

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/nathan-osman/go-herald"
	"github.com/nathan-osman/pratl/db"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const sessionName = "pratl"

// Server provides an HTTP interface for connecting clients.
type Server struct {
	server http.Server
	conn   *db.Conn
	herald *herald.Herald
	logger zerolog.Logger
}

// New creates a new server instance.
func New(cfg *Config) (*Server, error) {
	var (
		r = gin.Default()
		s = &Server{
			server: http.Server{
				Addr:    cfg.Addr,
				Handler: r,
			},
			conn:   cfg.Conn,
			herald: herald.New(),
			logger: log.With().Str("package", "server").Logger(),
		}
		store = cookie.NewStore([]byte(cfg.Key))
	)

	r.Use(sessions.Sessions(sessionName, store))

	// Unauthenticated methods
	r.POST("/auth/login", s.auth_login_POST)
	r.POST("/auth/register", s.auth_register_POST)

	// Authenticated methods
	a := r.Group("/").Use(s.requireLogin)

	a.POST("/rooms", s.rooms_POST)

	a.POST("/messages", s.messages_POST)

	a.GET("/ws", s.ws)

	// Setup and initialize the herald
	s.herald.MessageHandler = s.processMessage
	s.herald.ClientAddedHandler = s.processClientAdded
	s.herald.ClientRemovedHandler = s.processClientRemoved
	s.herald.Start()

	// Start the goroutine that listens for incoming connections
	go func() {
		defer s.logger.Info().Msg("server stopped")
		s.logger.Info().Msg("server started")
		if err := s.server.ListenAndServe(); errors.Is(err, http.ErrServerClosed) {
			s.logger.Error().Msg(err.Error())
		}
	}()

	return s, nil
}

// Close shuts down the server and herald.
func (s *Server) Close() {
	s.server.Shutdown(context.Background())
	s.herald.Close()
}
