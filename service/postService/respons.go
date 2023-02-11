package postservice

import "time"

type ImgUrl struct {
	ID  uint
	Url string
}

type Category struct {
	ID   uint
	Name string
}

type Post struct {
	ID            uint
	UserID        uint
	ImgUrls       []ImgUrl `gorm:"many2many:post_imgurls"`
	ImgWidthRate  uint
	ImgHeightRate uint
	Explain       string
	Categorys     []Category `gorm:"many2many:post_categories;unique"`
	LikesCount    uint
	CommentsCount uint
	Place         string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
