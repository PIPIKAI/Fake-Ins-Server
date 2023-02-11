package accountservice

import "time"

type UserDto struct {
	ID        uint
	UserName  string
	Telephone string
	Name      string
	Gender    string
	Email     string
	Photo     string
	Website   string
	Bio       string
	BirthDay  string

	WatchsCounts  uint
	FansCounts    uint
	CommentsCount uint
	PostsCounts   uint
	CreatedAt     time.Time
}
