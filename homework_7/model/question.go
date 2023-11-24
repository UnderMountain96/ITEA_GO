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

func NewQuestion(
	id uuid.UUID,
	text string,
	answerOptions map[uuid.UUID]string,
	correctAnswerID uuid.UUID,
) Question {
	return Question{
		id:              id,
		text:            text,
		answerOptions:   answerOptions,
		correctAnswerID: correctAnswerID,
	}
}
