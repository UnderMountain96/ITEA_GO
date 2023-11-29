package model

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type Article struct {
	Id        uuid.UUID
	Title     string
	Body      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewArticle(id uuid.UUID, title string) *Article {
	creationDate := time.Now()

	return &Article{
		Id:        id,
		Title:     title,
		CreatedAt: creationDate,
		UpdatedAt: creationDate,
	}
}

func (a *Article) Publish() error {
	if a.Body == "" {
		return errors.New("body is empty")
	}

	a.UpdatedAt = time.Now()

	return nil
}
