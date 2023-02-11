package model

import (
	"time"

	"gorm.io/gorm"
)

type ImgUrl struct {
	ID  uint `gorm:"primaryKey"`
	Url string
}
type Post struct {
	ID            uint `gorm:"primaryKey"`
	UserID        uint
	ImgUrls       []ImgUrl `gorm:"many2many:post_imgurls"`
	ImgWidthRate  uint
	ImgHeightRate uint
	Explain       string `gorm:"comment:自己的说明"`
	Comments      []Comment
	Categorys     []Category `gorm:"many2many:post_categories;unique"`
	Likes         []Like     `gorm:"polymorphic:Owner;unique"`
	DisLikes      []DisLike  `gorm:"unique;polymorphic:Owner;unique"`

	LikesCount    uint `gorm:"default:0"`
	CommentsCount uint `gorm:"default:0"`
	Place         string
	// IsLiked       bool `gorm:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (c *Post) AfterCreate(tx *gorm.DB) (err error) {
	var count int64
	tx.Model(Post{}).Where("user_id = ?", c.UserID).Count(&count)
	tx.Model(User{}).Where("ID = ?", c.UserID).UpdateColumn("posts_counts", count)
	return
}
func (c *Post) AfterDelete(tx *gorm.DB) (err error) {
	var count int64
	tx.Model(Post{}).Where("user_id = ?", c.UserID).Count(&count)
	tx.Model(User{}).Where("ID = ?", c.UserID).UpdateColumn("posts_counts", count)
	return
}

// func (u *Post) AfterFind(tx *gorm.DB) (err error) {
// 	tx.Model(u).Association("Likes")
// 	u.IsLiked :=
// 	return
// }
