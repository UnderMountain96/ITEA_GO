package store

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/UnderMountain96/ITEA_GO/student_testing_json/model"
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
)

type QuestionRepository struct {
	conn *sql.Conn
}

func NewQuestionRepository(conn *sql.Conn) *QuestionRepository {
	return &QuestionRepository{conn: conn}
}

func (r *QuestionRepository) Get(ctx context.Context, id uuid.UUID) ([]*model.Question, error) {

	rows, err := r.conn.QueryContext(ctx, "SELECT id, text, answers, correct_answer_id FROM view_question")
	if err != nil {
		return nil, fmt.Errorf("Get: query error: %w", err)
	}

	questions := make([]*model.Question, 0)

	for rows.Next() {
		var id, title, correct_answer_id string
		var answersJSON pgtype.JSONArray
		if err := rows.Scan(&id, &title, &answersJSON, &correct_answer_id); err != nil {
			return nil, fmt.Errorf("Get: scan error: %w", err)
		}

		answersMap := map[uuid.UUID]string{}
		for _, v := range answersJSON.Elements {
			if err := json.Unmarshal(v.Bytes, &answersMap); err != nil {
				return nil, fmt.Errorf("Get: scan error: %w", err)
			}
		}
		answers := []*model.Answer{}
		for i, t := range answersMap {
			a := model.Answer{
				ID:   i,
				Text: t,
			}
			answers = append(answers, &a)
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
