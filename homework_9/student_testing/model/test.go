package model

import "github.com/google/uuid"

type Test struct {
	ID             uuid.UUID `json:"id"`
	Title          string    `json:"title"`
	Questions      []*Question
	CorrectAnswers []uuid.UUID
}

func NewTest(id uuid.UUID, title string) *Test {
	return &Test{
		ID:    id,
		Title: title,
	}
}

func (t *Test) GetTest() *Test {
	return t
}

func (t *Test) SetQuestions(questions []*Question) {
	t.Questions = questions
}

func (t *Test) GetID() uuid.UUID {
	return t.ID
}

func (t *Test) GetTitle() string {
	return t.Title
}

func (t *Test) GetQuestions() []*Question {
	return t.Questions
}

func (t *Test) GetCorrectAnswerCount() int {
	return len(t.CorrectAnswers)
}

func (t *Test) GetWrongAnswerCount() int {
	return len(t.Questions) - t.GetCorrectAnswerCount()
}

func (t *Test) AddCorrectAnswer(id uuid.UUID) {
	t.CorrectAnswers = append(t.CorrectAnswers, id)
}
