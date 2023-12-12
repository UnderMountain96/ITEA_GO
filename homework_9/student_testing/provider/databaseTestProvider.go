package provider

import (
	"context"
	"fmt"

	"github.com/UnderMountain96/ITEA_GO/student_testing_json/model"
	"github.com/UnderMountain96/ITEA_GO/student_testing_json/store"
)

type DatabaseTestsProvider struct {
	testRepository     *store.TestRepository
	questionRepository *store.QuestionRepository
}

func NewDatabaseTestsProvider(testRepository *store.TestRepository, questionRepository *store.QuestionRepository) *DatabaseTestsProvider {
	return &DatabaseTestsProvider{testRepository: testRepository, questionRepository: questionRepository}
}

func (p *DatabaseTestsProvider) GetTests() ([]*model.Test, error) {
	ctx := context.Background()
	tests, err := p.testRepository.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("GetTests: cannot get tests: %w", err)
	}

	for _, test := range tests {
		questions, err := p.questionRepository.Get(ctx, test.ID)
		if err != nil {
			return nil, fmt.Errorf("GetTests: cannot get questions: %w", err)
		}

		test.Questions = questions
	}

	return tests, nil
}
