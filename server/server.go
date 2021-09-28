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
		router = gin.Default()
		s      = &Server{
			server: http.Server{
				Addr:    cfg.Addr,
				Handler: router,
			},
			conn:   cfg.Conn,
			herald: herald.New(),
			logger: log.With().Str("package", "server").Logger(),
		}
	)
	s.herald.MessageHandler = s.processMessage
	s.herald.Start()
	go func() {
		defer s.logger.Info().Msg("server stopped")
		s.logger.Info().Msg("server started")
		if err := s.server.ListenAndServe(); errors.Is(err, http.ErrServerClosed) {
			s.logger.Error().Msg(err.Error())
		}
	}()
	return s
}

// Close shuts down the server.
func (s *Server) Close() {
	s.server.Shutdown(context.Background())
}
