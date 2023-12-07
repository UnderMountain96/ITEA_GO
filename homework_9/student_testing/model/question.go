package model

import (
	"github.com/google/uuid"
)

type Question struct {
	ID              uuid.UUID `json:"id"`
	Text            string    `json:"text"`
	AnswerOptions   []*Answer `json:"answers"`
	CorrectAnswerID uuid.UUID `json:"correct_answer_id"`
}

type Answer struct {
	ID   uuid.UUID
	Text string
}

func NewQuestion(
	id uuid.UUID,
	text string,
	answerOptions []*Answer,
	correctAnswerID uuid.UUID,
) *Question {
	return &Question{
		ID:              id,
		Text:            text,
		AnswerOptions:   answerOptions,
		CorrectAnswerID: correctAnswerID,
	}
}

func (q *Question) IsCorrectAnswer(id uuid.UUID) bool {
	return id == q.CorrectAnswerID
}

func (t *Question) GetID() uuid.UUID {
	return t.ID
}

func (t *Question) GetText() string {
	return t.Text
}

func (t *Question) GetAnswerOptions() []*Answer {
	return t.AnswerOptions
}
