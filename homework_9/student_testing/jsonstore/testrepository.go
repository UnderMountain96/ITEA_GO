package jsonstore

import (
	"encoding/json"
	"fmt"

	"github.com/UnderMountain96/ITEA_GO/student_testing_json/model"
)

type TestRepository struct {
	data []byte
}

func NewTestRepository(data []byte) *TestRepository {
	return &TestRepository{data: data}
}

func (r *TestRepository) GetAll() ([]*model.Test, error) {
	availableTests := []*model.Test{}

	if err := json.Unmarshal(r.data, &availableTests); err != nil {
		return nil, fmt.Errorf("GetAll: query error: %w", err)
	}

	return availableTests, nil
}
