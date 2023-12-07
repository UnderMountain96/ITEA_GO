package provider

import (
	"fmt"

	"github.com/UnderMountain96/ITEA_GO/student_testing_json/jsonstore"
	"github.com/UnderMountain96/ITEA_GO/student_testing_json/model"
)

type JsonTestsProvider struct {
	testRepository     *jsonstore.TestRepository
	questionRepository *jsonstore.QuestionRepository
}

func NewJsonTestsProvider(testRepository *jsonstore.TestRepository, questionRepository *jsonstore.QuestionRepository) *JsonTestsProvider {
	return &JsonTestsProvider{testRepository: testRepository, questionRepository: questionRepository}
}

func (p *JsonTestsProvider) GetTests() ([]*model.Test, error) {
	tests, err := p.testRepository.GetAll()
	if err != nil {
		return nil, fmt.Errorf("GetTests: cannot get tests: %w", err)
	}

	for _, test := range tests {
		questions, err := p.questionRepository.Get(test.ID)
		if err != nil {
			return nil, fmt.Errorf("GetTests: cannot get questions: %w", err)
		}

		test.Questions = questions
	}

	return tests, nil
}
