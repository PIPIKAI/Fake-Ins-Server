package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Role      string `gorm:"default:user"`
	UserName  string `gorm:"not null;unique;"`
	Telephone string `gorm:"size:11;comment:手机号"`
	Name      string `gorm:"not null;comment:姓名"`
	Gender    string
	Email     string `gorm:"comment:邮箱"`
	Photo     string
	Website   string
	Bio       string
	BirthDay  string

	PassWord string `gorm:"not null;"`

	Watchs      []*User      `gorm:"many2many:user_watchs;comment:关注"`
	Fans        []*User      `gorm:"many2many:user_fans;comment:粉丝"`
	Comments    []Comment    `gorm:"comment:评论"`
	Collections []Collection `gorm:"comment:收藏"`
	Posts       []Post

	WatchsCounts  uint `gorm:"default:0"`
	FansCounts    uint `gorm:"default:0"`
	CommentsCount uint `gorm:"default:0"`
	PostsCounts   uint `gorm:"default:0"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	if u.ID == 1 {
		tx.Model(u).Update("role", "admin")
	}
	// 创建头像
	if u.Photo == "" {
		tx.Model(u).Update("photo", "http://pic.kiass.top/1673942041153account.png")
	}
	return
}
