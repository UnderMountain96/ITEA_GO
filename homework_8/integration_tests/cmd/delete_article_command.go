package cmd

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/greeflas/itea_golang/repository"
)

type DeleteArticleCommand struct {
	articleRepository *repository.ArticleRepository
}

func NewDeleteArticleCommand(articleRepository *repository.ArticleRepository) *DeleteArticleCommand {
	return &DeleteArticleCommand{articleRepository: articleRepository}
}

func (c *DeleteArticleCommand) Name() string {
	return "delete_article"
}

func (c *DeleteArticleCommand) Run(ctx context.Context, params map[string]string) error {
	idStr, ok := params["id"]
	if !ok {
		return errors.New("error: id param is required for update")
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		return err
	}

	a, err := c.articleRepository.Get(ctx, id)
	if err != nil {
		return err
	}

	if err := c.articleRepository.Delete(ctx, a.Id); err != nil {
		return err
	}

	return nil
}
