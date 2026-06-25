package entity

type Comment struct {
	ID      int64
	UserID  int64
	PostID  int64
	Content string
}
