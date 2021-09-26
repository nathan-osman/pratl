package db

// Config defines the parameters for creating Conn instances.
type Config struct {
	Host     string
	Port     int
	Name     string
	User     string
	Password string
}
