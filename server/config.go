package server

import (
	"github.com/nathan-osman/pratl/db"
)

// Config defines the parameters for creating server instances.
type Config struct {
	Addr string
	Conn *db.Conn
}
