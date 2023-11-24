package model

import (
	"github.com/google/uuid"
)

type Question struct {
<<<<<<< HEAD
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
=======
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
>>>>>>> 2e2796c (add sql store)
	}
}
