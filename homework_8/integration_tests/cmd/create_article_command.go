package cmd

import (
	"context"
	"github.com/google/uuid"
	"github.com/greeflas/itea_golang/model"
	"github.com/greeflas/itea_golang/repository"
)

type CreateArticleCommand struct {
	articleRepository *repository.ArticleRepository
}

func NewCreateArticleCommand(articleRepository *repository.ArticleRepository) *CreateArticleCommand {
	return &CreateArticleCommand{articleRepository: articleRepository}
}

func (c *CreateArticleCommand) Name() string {
	return "create_article"
}

func (c *CreateArticleCommand) Run(ctx context.Context) error {
	article := model.NewArticle(uuid.New(), "This is title!")

	return c.articleRepository.Create(ctx, article)
}
