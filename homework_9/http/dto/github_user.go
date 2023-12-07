package dto

import "time"

type GitHubUser struct {
	Id        int       `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Bio       string    `json:"bio,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
