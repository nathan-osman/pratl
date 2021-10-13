package db

import (
	"time"
)

// Message represents an individual message.
type Message struct {
	ID           int64     `json:"id"`
	MemberID     int64     `gorm:"not null" json:"-"`
	Member       *Member   `gorm:"constraint:OnDelete:CASCADE" json:"member"`
	Body         string    `gorm:"not null" json:"body"`
	CreationDate time.Time `gorm:"not null" json:"creation_date"`
	StarCount    int64     `gorm:"not null" json:"star_count"`
	Stars        []*User   `gorm:"many2many:message_stars;" json:"-"`
	IsEdited     bool      `gorm:"not null" json:"is_edited"`
}
