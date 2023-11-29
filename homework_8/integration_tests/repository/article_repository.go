package repository

import (
	"context"
	"github.com/greeflas/itea_golang/model"
	"github.com/jackc/pgx/v5"
)

type ArticleRepository struct {
	conn *pgx.Conn
}

func NewArticleRepository(conn *pgx.Conn) *ArticleRepository {
	return &ArticleRepository{conn: conn}
}

func (r *ArticleRepository) Create(ctx context.Context, a *model.Article) error {
	sql := `
INSERT INTO articles (id, title, body, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5)
`

	_, err := r.conn.Exec(
		ctx,
		sql,
		a.Id,
		a.Title,
		a.Body,
		a.CreatedAt,
		a.UpdatedAt,
	)

	return err
}
