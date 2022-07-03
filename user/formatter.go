package user

import "time"

type UserFormatter struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
}

type CheckEmailFormatter struct {
	IsAvailable bool `json:"is_available"`
}

func FormatUser(user User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Token:     token,
		CreatedAt: user.CreatedAt,
	}
	return formatter
}

func FormatCheckEmail(isEmailAvailable bool) CheckEmailFormatter {
	formatter := CheckEmailFormatter{
		IsAvailable: isEmailAvailable,
	}

	return formatter
}
