package server

import (
	"github.com/nathan-osman/pratl/db"
)

// Config defines the parameters for creating server instances.
type Config struct {
	Debug bool
	Addr  string
	Key   string
	Conn  *db.Conn
}
