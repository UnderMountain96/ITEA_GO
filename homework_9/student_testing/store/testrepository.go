package store

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/UnderMountain96/ITEA_GO/student_testing_json/model"
	"github.com/google/uuid"
)

type TestRepository struct {
	conn *sql.Conn
}

func NewTestRepository(conn *sql.Conn) *TestRepository {
	return &TestRepository{conn: conn}
}
func (r *TestRepository) GetAll(ctx context.Context) ([]*model.Test, error) {
	rows, err := r.conn.QueryContext(ctx, "SELECT id, title FROM test")
	if err != nil {
		return nil, fmt.Errorf("GetAll: query error: %w", err)
	}

	availableTests := make([]*model.Test, 0)

	for rows.Next() {
		var id, title string
		if err := rows.Scan(&id, &title); err != nil {
			fmt.Printf("GetAll: scan error: %s\n", err)
		}
		t := model.NewTest(uuid.MustParse(id), title)

		availableTests = append(availableTests, t)
	}

	return availableTests, nil
}
