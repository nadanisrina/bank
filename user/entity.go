package user

import "time"

type User struct {
	ID           int
	Username     string
	Email        string
	PasswordHash string
	Token        string
	CreatedAt    time.Time
}
