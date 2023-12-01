package cmd

import (
	"context"

	"github.com/greeflas/itea_golang/repository"
)

type GetAllArticleCommand struct {
	articleRepository *repository.ArticleRepository
}

func NewGetAllArticleCommand(articleRepository *repository.ArticleRepository) *GetAllArticleCommand {
	return &GetAllArticleCommand{articleRepository: articleRepository}
}

func (c *GetAllArticleCommand) Name() string {
	return "get_all_article"
}

func (c *GetAllArticleCommand) Run(ctx context.Context, params map[string]string) error {
	articals, err := c.articleRepository.GetAll(ctx)
	if err != nil {
		return err
	}

	for _, a := range articals {
		a.ShowInfo()
	}
	return nil
}
