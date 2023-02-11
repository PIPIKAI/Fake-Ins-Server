package model

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID           uint `gorm:"primaryKey"`
	UserID       uint
	PostID       uint
	Content      string
	ReplyID      *uint
	Replys       []Comment `gorm:"foreignkey:ReplyID"`
	ReplysCounts uint      `gorm:"default:0"`
	Likes        []Like    `gorm:"unique;polymorphic:Owner;"`
	DisLikes     []DisLike `gorm:"unique;polymorphic:Owner;"`
	LikesCounts  uint      `gorm:"default:0"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (c *Comment) AfterCreate(tx *gorm.DB) (err error) {
	var ucounts, pcounts, replycounts int64
	if c.ReplyID != nil {
		tx.Model(Comment{}).Where("reply_id = ?", c.ReplyID).Count(&replycounts)
		tx.Model(Comment{ID: *c.ReplyID}).UpdateColumn("replys_counts", replycounts)
	}
	tx.Model(Comment{}).Where("post_id = ? AND reply_id is NULL", c.PostID).Count(&pcounts)
	tx.Model(Comment{}).Where("user_id = ?", c.UserID).Count(&ucounts)
	tx.Model(User{ID: c.UserID}).UpdateColumn("comments_count", ucounts)
	tx.Model(Post{ID: c.PostID}).UpdateColumn("comments_count", pcounts)
	return
}
func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	var ucounts, pcounts int64
	tx.Model(Comment{}).Where("post_id = ?", c.PostID).Count(&pcounts)
	tx.Model(Comment{}).Where("user_id = ?", c.UserID).Count(&ucounts)
	tx.Model(User{ID: c.UserID}).UpdateColumn("comments_count", ucounts)
	tx.Model(Post{ID: c.PostID}).UpdateColumn("comments_count", pcounts)
	return
}
