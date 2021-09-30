package server

import (
	"context"
	"errors"
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/nathan-osman/go-herald"
	"github.com/nathan-osman/pratl/db"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const identityKey = "user"

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
	)

	// Setup the JWT middleware
	m, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:           "Pratl",
		Key:             []byte(cfg.Key),
		Authenticator:   s.authenticator,
		Authorizator:    s.authorizator,
		PayloadFunc:     s.payloadFunc,
		Unauthorized:    e,
		IdentityHandler: s.identityHandler,
		IdentityKey:     identityKey,
	})
	if err != nil {
		return nil, err
	}

	// JWT authentication methods
	r.POST("/auth/login", m.LoginHandler)
	r.POST("/auth/refresh", m.RefreshHandler)
	r.POST("/auth/register", s.auth_register_POST)

	a := r.Group("/").Use(m.MiddlewareFunc())

	// Protected API methods

	a.GET("/rooms", s.rooms_GET)
	a.POST("/rooms", s.rooms_POST)
	a.PUT("/rooms/:id", s.rooms_id_PUT)
	a.DELETE("/rooms/:id", s.rooms_id_DELETE)

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
