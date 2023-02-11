package userservice

// type ResUser struct {
// 	ID        uint
// 	UserName  string
// 	Telephone string
// 	Name      string
// 	Gender    string
// 	Email     string
// 	Photo     string
// 	Website   string
// 	Bio       string
// 	BirthDay  string
// }
type RESUser struct {
	ID           uint
	UserName     string
	Name         string
	Photo        string
	Website      string
	Bio          string
	PostsCounts  uint
	WatchsCounts uint
	FansCounts   uint
	IsWatched    bool `gorm:"-"`
}
