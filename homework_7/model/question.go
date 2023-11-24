package model

import (
	"github.com/google/uuid"
)

type Question struct {
	Text          string
	AnswerOptions map[uuid.UUID]string
	CorrectAnswer uuid.UUID
}

func (q *Question) IsCorrectAnswer(ca uuid.UUID) bool {
	return ca == q.CorrectAnswer
}

func NewQuestion(text string, answerOptions map[uuid.UUID]string, correctAnswer uuid.UUID) Question {
	return Question{
		Text:          text,
		AnswerOptions: answerOptions,
		CorrectAnswer: correctAnswer,
	}
}
