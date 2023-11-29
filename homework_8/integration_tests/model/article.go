package model

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
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

func (a *Article) ShowInfo() {
	fmt.Printf("Id: \t\t%s\n", a.Id)
	fmt.Printf("Title: \t\t%s\n", a.Title)
	fmt.Printf("CreatedAt: \t%s\n", a.CreatedAt.Format("02.01.2006 15:04:05"))
	fmt.Printf("UpdatedAt: \t%s\n", a.UpdatedAt.Format("02.01.2006 15:04:05"))
	fmt.Println()
}
