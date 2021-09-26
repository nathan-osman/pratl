package server

import (
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nathan-osman/pratl/db"
)

// Server provides an HTTP interface for connecting clients.
type Server struct {
	listener net.Listener
	conn     *db.Conn
	stopped  chan bool
}

// New creates a new server instance.
func New(cfg *Config) (*Server, error) {
	l, err := net.Listen("tcp", cfg.Addr)
	if err != nil {
		return nil, err
	}
	var (
		s = &Server{
			listener: l,
			conn:     cfg.Conn,
			stopped:  make(chan bool),
		}
		router = mux.NewRouter()
		server = http.Server{
			Handler: router,
		}
	)
	go func() {
		defer close(s.stopped)
		if err := server.Serve(l); err != http.ErrServerClosed {
			// TODO: display error
		}
	}()
	return s, err
}

// Close shuts down the server.
func (s *Server) Close() {
	s.listener.Close()
	<-s.stopped
}
