package entities

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        string `gorm:"primary_key" json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	IsAdmin   bool   `json:"isAdmin"`
	CreatedAt string `db:"created_at" json:"created_at"`
}

type UserResponse struct {
	ID        string `gorm:"primary_key" json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	IsAdmin   bool   `json:"isAdmin"`
	CreatedAt string `db:"created_at" json:"created_at"`
}

func NewUser() User {
	return User{ID: uuid.NewString(), CreatedAt: time.Now().Format("2006-01-02 15:04:05")}
}
