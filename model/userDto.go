package model

import "time"

type UserDto struct {
	ID        uint   `gorm:"primaryKey"`
	UserName  string `gorm:"not null;unique;"`
	Telephone string `gorm:"size:11;comment:手机号"`
	Name      string `gorm:"not null;comment:姓名"`
	Gender    string
	Email     string `gorm:"comment:邮箱"`
	Photo     string
	Website   string
	Bio       string
	BirthDay  string

	WatchsCounts   uint
	FansCounts     uint
	CommentsCounts uint
	PostsCounts    uint
	CreatedAt      time.Time
}

type OtherUserDto struct {
	ID       uint   `gorm:"primaryKey"`
	UserName string `gorm:"not null;unique;"`
	Name     string `gorm:"not null;comment:姓名"`
	Photo    string
	Website  string
	Bio      string

	WatchsCounts uint
	FansCounts   uint
	PostsCounts  uint
	CreatedAt    time.Time
}
