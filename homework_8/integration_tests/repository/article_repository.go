package repository

import (
	"context"
	"time"

	"github.com/UnderMountain96/ITEA_GO/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type ArticleRepository struct {
	conn *pgx.Conn
}

func NewArticleRepository(conn *pgx.Conn) *ArticleRepository {
	return &ArticleRepository{conn: conn}
}

func (r *ArticleRepository) Create(ctx context.Context, a *model.Article) error {
	sql := `INSERT INTO articles (id, title, body, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)`

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

func (r *ArticleRepository) Get(ctx context.Context, id uuid.UUID) (*model.Article, error) {
	sql := `SELECT id, title, body, created_at, updated_at FROM articles WHERE id = $1`

	row := r.conn.QueryRow(
		ctx,
		sql,
		id,
	)

	a := &model.Article{}

	if err := row.Scan(
		&a.Id,
		&a.Title,
		&a.Body,
		&a.CreatedAt,
		&a.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return a, nil
}

func (r *ArticleRepository) Update(ctx context.Context, a *model.Article) error {
	sql := `UPDATE articles SET (title, body, updated_at) = ($2, $3, $4) WHERE id = $1`

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
