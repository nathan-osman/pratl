package db

import (
	"fmt"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Conn maintains a connection to the database.
type Conn struct {
	*gorm.DB
	logger zerolog.Logger
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
	return &Conn{
		DB:     d,
		logger: log.With().Str("package", "db").Logger(),
	}, nil
}

// Migrate applies all pending database migrations.
func (c *Conn) Migrate() error {
	c.logger.Info().Msg("applying migrations...")
	return c.AutoMigrate(
		&User{},
		&Room{},
		&Member{},
		&Message{},
	)
}

// Transaction runs the specified function in a transaction.
func (c *Conn) Transaction(fn func(*Conn) error) error {
	return c.DB.Transaction(func(tx *gorm.DB) error {
		return fn(&Conn{
			DB:     tx,
			logger: c.logger,
		})
	})
}

// Close closes the database connection.
func (c *Conn) Close() {
	db, _ := c.DB.DB()
	db.Close()
}
