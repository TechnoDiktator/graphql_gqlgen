package entity

import "time"

type Comment struct {
	ID        int64
	UserID    int64
	PostID    int64
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
