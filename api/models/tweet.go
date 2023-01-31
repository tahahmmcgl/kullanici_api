package models

import (
	"html"
	"strings"
	"time"
)

type Tweet struct {
	ID        uint      `gorm:"primary key:auto_increment"json:"id"`
	OwnerID   uint      `gorm:"not null;"json:"owner_id"`
	Content   string    `gorm:"type:varchar(100);not null"json:"content"`
	CreatedAt time.Time `gorm:"default:'0000-00-00 00:00:00'"json:"created_at"`
	UpdatedAt time.Time `gorm:"default:'0000-00-00 00:00:00'"json:"updated_at"`
}

func (t *Tweet) Prepare() {
	t.ID = 0
	t.OwnerID = 0
	t.Content = html.EscapeString(strings.TrimSpace(t.Content))
	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()
}
