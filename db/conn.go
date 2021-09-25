package db

import (
	"gorm.io/gorm"
)

// Conn maintains a connection to the database.
type Conn struct {
	*gorm.DB
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
