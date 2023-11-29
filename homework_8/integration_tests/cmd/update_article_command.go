package cmd

import (
	"context"
	"errors"
	"flag"
	"os"

	"github.com/google/uuid"
	"github.com/greeflas/itea_golang/model"
	"github.com/greeflas/itea_golang/repository"
)

const (
	usageTitle = "New actical title"
	usageBody  = "New actical title"
)

type UpdateArticleCommand struct {
	articleRepository *repository.ArticleRepository
}

func NewUpdateArticleCommand(articleRepository *repository.ArticleRepository) *UpdateArticleCommand {
	return &UpdateArticleCommand{articleRepository: articleRepository}
}

func (c *UpdateArticleCommand) Name() string {
	return "update_article"
}

func (c *UpdateArticleCommand) Run(ctx context.Context) error {
	cmd := flag.NewFlagSet("update", flag.ExitOnError)
	var idStr, title, body string

	cmd.StringVar(&idStr, "id", "", "Actical ID")
	cmd.StringVar(&title, "title", "", "New actical title")
	cmd.StringVar(&title, "t", "", usageTitle+" (shorthand)")
	cmd.StringVar(&body, "body", "", "New actical body")
	cmd.StringVar(&body, "b", "", usageBody+" (shorthand)")

	if err := cmd.Parse(os.Args[2:]); err != nil {
		return nil
	}

	if idStr == "" {
		return errors.New("id must not be empty")
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		return err
	}

	a := &model.Article{
		Id: id,
	}
	if err := c.articleRepository.Get(ctx, a); err != nil {
		return err
	}

	if title != "" {
		a.Title = title
	}

	if body != "" {
		a.Body = body
	}

	if err := c.articleRepository.Update(ctx, a); err != nil {
		return err
	}

	return nil
}
