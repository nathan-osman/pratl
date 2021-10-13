package db

// Member represents a member of a specific room.
type Member struct {
	ID      int64 `json:"id"`
	UserID  int64 `gorm:"not null" json:"-"`
	User    *User `gorm:"constraint:OnDelete:CASCADE" json:"user"`
	RoomID  int64 `gorm:"not null" json:"-"`
	Room    *Room `gorm:"constraint:OnDelete:CASCADE" json:"room"`
	IsOwner bool  `gorm:"not null" json:"is_owner"`
	IsAdmin bool  `gorm:"not null" json:"is_admin"`
}
