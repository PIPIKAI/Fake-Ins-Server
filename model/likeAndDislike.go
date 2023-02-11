package model

import (
	"time"

	"gorm.io/gorm"
)

type Like struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"primaryKey;autoIncrement:false"`
	OwnerID   uint   `gorm:"primaryKey;autoIncrement:false"`
	OwnerType string `gorm:"primaryKey;autoIncrement:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
type DisLike struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	OwnerID   uint
	OwnerType string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (c *Like) AfterCreate(tx *gorm.DB) (err error) {
	// 帖子
	var counts int64
	if c.OwnerType == "posts" {
		counts = tx.Model(Post{ID: c.OwnerID}).Association("Likes").Count()
		tx.Model(&Post{ID: c.OwnerID}).UpdateColumn("likes_count", counts)
	} else if c.OwnerType == "comments" {
		counts = tx.Model(Comment{ID: c.OwnerID}).Association("Likes").Count()
		tx.Model(&Comment{ID: c.OwnerID}).UpdateColumn("likes_counts", counts)
	}
	return
}
func (c *Like) AfterDelete(tx *gorm.DB) (err error) {
	var counts int64
	if c.OwnerType == "posts" {
		counts = tx.Model(Post{ID: c.OwnerID}).Association("Likes").Count()
		tx.Model(&Post{ID: c.OwnerID}).UpdateColumn("likes_count", counts)
	} else if c.OwnerType == "comments" {
		counts = tx.Model(Comment{ID: c.OwnerID}).Association("Likes").Count()
		tx.Model(&Comment{ID: c.OwnerID}).UpdateColumn("likes_counts", counts)
	}
	return
}
