package payload

import (
	"time"
)

type (
	User struct {
		CreatedAt   time.Time `json:"created_at"`
		DisplayName string    `json:"display_name"`
		Email       string    `json:"email"`
	}
	UserList  map[int64]User
	UserStore struct {
		Increment int64    `json:"increment"`
		List      UserList `json:"list"`
	}
	CreateUserRequest struct {
		DisplayName string `json:"display_name"`
		Email       string `json:"email"`
	}
	UpdateUserRequest struct {
		DisplayName string `json:"display_name"`
		Email       string `json:"email"`
	}
)

func CreateUser(name, email string) User {
	return User{
		CreatedAt:   time.Now(),
		DisplayName: name,
		Email:       email,
	}
}
