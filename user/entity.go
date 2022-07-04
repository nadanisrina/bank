package user

import "time"

type User struct {
	ID             int
	Username       string
	Email          string
	PasswordHash   string
	Token          string
	AvatarFileName string
	CreatedAt      time.Time
}
