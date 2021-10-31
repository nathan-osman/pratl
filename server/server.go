package server

import (
	"context"
	"errors"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/nathan-osman/go-herald"
	"github.com/nathan-osman/pratl/db"
	"github.com/nathan-osman/pratl/ui"
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

type embedFileSystem struct {
	http.FileSystem
}

func (e embedFileSystem) Exists(prefix, path string) bool {
	_, err := e.Open(path)
	return !os.IsNotExist(err)
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

	r.Use(
		sessions.Sessions(sessionName, store),
		static.Serve("/", embedFileSystem{FileSystem: http.FS(ui.Content)}),
	)

	// If operating in debug mode, add CORS headers
	if cfg.Debug {
		r.Use(cors.New(cors.Config{
			AllowOrigins: []string{"*"},
			AllowHeaders: []string{"Content-Type"},
		}))
	}

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
