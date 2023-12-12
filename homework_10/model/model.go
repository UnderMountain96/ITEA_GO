package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	UpdatedAt time.Time `json:"update_at"`
	CreatedAt time.Time `json:"create_at"`
}

func (u *User) Update(username string, email string) {
	u.Username = username
	u.Email = email
	u.UpdatedAt = time.Now()
}

func NewUser(username string, email string) *User {
	timeNow := time.Now()
	return &User{
		ID:        uuid.New(),
		Username:  username,
		Email:     email,
		UpdatedAt: timeNow,
		CreatedAt: timeNow,
	}
}
