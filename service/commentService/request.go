package commentservice

type ReqComment struct {
	UserID  uint
	PostID  uint
	Content string
	// Replys       []Comment `gorm:"foreignkey:ReplyID"`
}
type ReqReply struct {
	UserID  uint
	PostID  uint
	ReplyID uint
	Content string
	// Replys       []Comment `gorm:"foreignkey:ReplyID"`
}
