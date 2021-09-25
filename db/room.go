package db

import (
	"time"
)

// Room represents a chat group.
type Room struct {
	ID           int64     `json:"id"`
	Name         string    `gorm:"not null" json:"name"`
	CreationDate time.Time `gorm:"not null" json:"creation_date"`
}
