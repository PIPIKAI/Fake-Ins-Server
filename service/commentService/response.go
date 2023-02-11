package commentservice

import "time"

type Comment struct {
	ID           uint
	UserID       uint
	PostID       uint
	Content      string
	ReplyID      *uint
	Replys       []Comment `gorm:"foreignkey:ReplyID"`
	ReplysCounts uint
	LikesCounts  uint
	CreatedAt    time.Time
}
