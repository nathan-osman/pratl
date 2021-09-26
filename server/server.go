package server

import (
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nathan-osman/go-herald"
	"github.com/nathan-osman/pratl/db"
)

// Server provides an HTTP interface for connecting clients.
type Server struct {
	listener net.Listener
	conn     *db.Conn
	herald   *herald.Herald
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
			herald:   herald.New(),
			stopped:  make(chan bool),
		}
		router = mux.NewRouter()
		server = http.Server{
			Handler: router,
		}
	)
	// TODO: initialize the herald
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
