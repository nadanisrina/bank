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
type UploadFileFormatter struct {
	IsUploaded bool `json:"is_uploaded"`
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

func FormatUploadAvatar(isFileUploaded bool) UploadFileFormatter {
	formatter := UploadFileFormatter{
		IsUploaded: isFileUploaded,
	}

	return formatter
}
