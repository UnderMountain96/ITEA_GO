package cmd

import (
	"context"

	"github.com/UnderMountain96/ITEA_GO/model"
	"github.com/UnderMountain96/ITEA_GO/repository"
	"github.com/google/uuid"
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

func (c *CreateArticleCommand) Run(ctx context.Context, params map[string]string) error {
	article := model.NewArticle(uuid.New(), "This is title!")

	return c.articleRepository.Create(ctx, article)
}
