package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Conn maintains a connection to the database.
type Conn struct {
	*gorm.DB
}

// New attempts to connect to the database.
func New(cfg *Config) (*Conn, error) {
	d, err := gorm.Open(
		postgres.Open(
			fmt.Sprintf(
				"host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
				cfg.Host,
				cfg.Port,
				cfg.Name,
				cfg.User,
				cfg.Password,
			),
		),
	)
	if err != nil {
		return nil, err
	}
	return &Conn{DB: d}, nil
}

// Migrate applies all pending database migrations.
func (c *Conn) Migrate() error {
	return c.AutoMigrate(
		&User{},
		&Room{},
		&Member{},
		&Message{},
	)
}

// Close closes the database connection.
func (c *Conn) Close() {
	db, _ := c.DB.DB()
	db.Close()
}
