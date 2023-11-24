package model

import "github.com/google/uuid"

type Test struct {
	id             uuid.UUID
	title          string
	questions      []Question
	correctAnswers []uuid.UUID
}

<<<<<<< HEAD
func NewTest(id uuid.UUID, title string, questions []Question) Test {
	// TODO: fetch real questions from API
	return Test{
		id:        id,
		title:     title,
		questions: questions,
	}
}

=======
func NewTest(id uuid.UUID, title string) Test {
	// TODO: fetch real questions from API
	return Test{
		id:    id,
		title: title,
	}
}

func (t *Test) SetQuestions(questions []Question) {
	t.questions = questions
}

>>>>>>> 2e2796c (add sql store)
func (t *Test) GetID() uuid.UUID {
	return t.id
}

func (t *Test) GetTitle() string {
	return t.title
}

func (t *Test) GetQuestions() []Question {
	return t.questions
}

func (t *Test) GetCorrectAnswerCount() int {
	return len(t.correctAnswers)
}

func (t *Test) GetWrongAnswerCount() int {
	return len(t.questions) - t.GetCorrectAnswerCount()
}

func (t *Test) AddCorrectAnswer(id uuid.UUID) {
	t.correctAnswers = append(t.correctAnswers, id)
}
