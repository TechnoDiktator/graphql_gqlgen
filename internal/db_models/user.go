package entity

import "time"

type User struct {
	ID           int64
	Name         string
	Email        string
	PasswordHash string
	Age          int32
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
