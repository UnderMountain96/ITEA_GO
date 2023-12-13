package jsonstore

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/UnderMountain96/ITEA_GO/student_testing_json/model"
	"github.com/google/uuid"
)

type QuestionRepository struct {
	data []byte
}

func NewQuestionRepository(data []byte) *QuestionRepository {
	return &QuestionRepository{data: data}
}

func (r *QuestionRepository) Get(id uuid.UUID) ([]*model.Question, error) {
	questions := map[string][]*model.Question{}

	if err := json.Unmarshal(r.data, &questions); err != nil {
		return nil, fmt.Errorf("Get: query error: %w", err)
	}

	q, ok := questions[id.String()]

	if !ok {
		return nil, errors.New("Get: test ton found")
	}

	return q, nil
}
