package model

import "time"

type Collection struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	Name      string `gorm:"default:default;unique"`
	Posts     []Post `gorm:"many2many:collection_posts;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
