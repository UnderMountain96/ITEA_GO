package cmd

import (
	"context"
	"flag"
	"os"

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

func (c *DeleteArticleCommand) Run(ctx context.Context) error {
	cmd := flag.NewFlagSet("delete", flag.ExitOnError)
	var idStr string

	cmd.StringVar(&idStr, "id", "", "Actical ID")

	if err := cmd.Parse(os.Args[2:]); err != nil {
		return nil
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
