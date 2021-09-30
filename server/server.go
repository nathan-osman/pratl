package server

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nathan-osman/go-herald"
	"github.com/nathan-osman/pratl/db"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Server provides an HTTP interface for connecting clients.
type Server struct {
	server http.Server
	conn   *db.Conn
	herald *herald.Herald
	logger zerolog.Logger
}

// New creates a new server instance.
func New(cfg *Config) *Server {
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
	)

	// Register the API routes

	r.POST("/users/login", s.users_login_POST)
	r.POST("/users/register", s.users_register_POST)

	r.GET("/rooms", s.rooms_GET)
	r.POST("/rooms", s.rooms_POST)
	r.PUT("/rooms/:id", s.rooms_id_PUT)
	r.DELETE("/rooms/:id", s.rooms_id_DELETE)

	r.POST("/messages", s.messages_POST)

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

	return s
}

// Close shuts down the server and herald.
func (s *Server) Close() {
	s.server.Shutdown(context.Background())
	s.herald.Close()
}
