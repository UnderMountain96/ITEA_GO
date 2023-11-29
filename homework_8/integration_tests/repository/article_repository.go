package repository

import (
	"context"
	"time"

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

func (r *ArticleRepository) Get(ctx context.Context, a *model.Article) error {
	sql := `SELECT title, body, created_at, updated_at FROM articles WHERE id = $1`

	row := r.conn.QueryRow(
		ctx,
		sql,
		a.Id,
	)

	if err := row.Scan(
		&a.Title,
		&a.Body,
		&a.CreatedAt,
		&a.UpdatedAt,
	); err != nil {
		return err
	}

	return nil
}

func (r *ArticleRepository) GetAll(ctx context.Context) ([]*model.Article, error) {
	sql := `SELECT id, title, body, created_at, updated_at FROM articles`

	rows, err := r.conn.Query(
		ctx,
		sql,
	)

	if err != nil {
		return nil, err
	}

	acticals := make([]*model.Article, 0)

	for rows.Next() {
		a := &model.Article{}
		rows.Scan(
			&a.Id,
			&a.Title,
			&a.Body,
			&a.CreatedAt,
			&a.UpdatedAt,
		)
		acticals = append(acticals, a)
	}

	return acticals, err
}

func (r *ArticleRepository) Update(ctx context.Context, a *model.Article) error {
	sql := `UPDATE articles SET title = $2, body = $3, updated_at = $4 WHERE id = $1`

	_, err := r.conn.Exec(
		ctx,
		sql,
		a.Id,
		a.Title,
		a.Body,
		time.Now(),
	)

	return err
}
