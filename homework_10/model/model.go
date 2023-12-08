package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	UpdateAt time.Time `json:"update_at"`
	CreateAt time.Time `json:"create_at"`
}

func NewUser(username string, email string) *User {
	return &User{
		ID:       uuid.New(),
		Username: username,
		Email:    email,
		UpdateAt: time.Now(),
		CreateAt: time.Now(),
	}
}
