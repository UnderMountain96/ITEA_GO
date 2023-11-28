package model

import (
	"github.com/google/uuid"
)

type Question struct {
	id              uuid.UUID
	text            string
	answerOptions   map[uuid.UUID]string
	correctAnswerID uuid.UUID
}

func (q *Question) IsCorrectAnswer(id uuid.UUID) bool {
	return id == q.correctAnswerID
}

func (t *Question) GetID() uuid.UUID {
	return t.id
}

func (t *Question) GetText() string {
	return t.text
}

func (t *Question) GetAnswerOptions() map[uuid.UUID]string {
	return t.answerOptions
}

func NewQuestion(
	id uuid.UUID,
	text string,
	answerOptions map[uuid.UUID]string,
	correctAnswerID uuid.UUID,
) *Question {
	return &Question{
		id:              id,
		text:            text,
		answerOptions:   answerOptions,
		correctAnswerID: correctAnswerID,
	}
}
