package cmd

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/greeflas/itea_golang/repository"
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

func (c *UpdateArticleCommand) Run(ctx context.Context, params map[string]string) error {
	idStr, ok := params["id"]
	if !ok {
		return errors.New("error: id param is required for update")
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		return err
	}

	title := params["title"]
	body := params["body"]

	a, err := c.articleRepository.Get(ctx, id)
	if err != nil {
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
