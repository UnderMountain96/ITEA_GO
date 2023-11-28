package sqlstore

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/UnderMountain96/ITEA_GO/student_testing/model"
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
)

type TestRepository struct {
	store *Store
}

func (r *TestRepository) GetAll(ctx context.Context) ([]*model.Test, error) {
	conn, err := r.store.db.Conn(ctx)
	if err != nil {
		return nil, fmt.Errorf("GetAll: connect error: %w", err)
	}
	defer conn.Close()

	rows, err := conn.QueryContext(ctx, "SELECT id, title FROM test")
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

func (r *TestRepository) GetQuestions(ctx context.Context, id uuid.UUID) ([]*model.Question, error) {
	conn, err := r.store.db.Conn(ctx)
	if err != nil {
		return nil, fmt.Errorf("GetQuestions: connect error: %w", err)
	}
	defer conn.Close()

	rows, err := conn.QueryContext(ctx, "SELECT id, text, answers, correct_answer_id FROM view_question")
	if err != nil {
		return nil, fmt.Errorf("GetQuestions: query error: %w", err)
	}

	questions := make([]*model.Question, 0)

	for rows.Next() {
		var id, title, correct_answer_id string
		var answersJSON pgtype.JSONArray
		if err := rows.Scan(&id, &title, &answersJSON, &correct_answer_id); err != nil {
			return nil, fmt.Errorf("GetQuestions: scan error: %w", err)
		}

		answers := map[uuid.UUID]string{}
		for _, v := range answersJSON.Elements {
			if err := json.Unmarshal(v.Bytes, &answers); err != nil {
				return nil, fmt.Errorf("GetQuestions: scan error: %w", err)
			}
		}

		questions = append(
			questions,
			model.NewQuestion(
				uuid.MustParse(id),
				title,
				answers,
				uuid.MustParse(correct_answer_id),
			),
		)
	}

	return questions, nil
}
