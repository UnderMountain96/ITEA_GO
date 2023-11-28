package main

import (
	"fmt"

	"github.com/UnderMountain96/ITEA_GO/student_testing/model"
	"github.com/google/uuid"
)

func showAvailableTests(tests []*model.Test) (*model.Test, error) {
	for idx, t := range tests {
		fmt.Printf("%d) %s.\n", idx+1, t.GetTitle())
	}

	fmt.Printf("Choose test: ")
	var testNum int
	_, err := fmt.Scan(&testNum)
	if err != nil {
		return &model.Test{}, fmt.Errorf("showAvailableTests: invalid number test value entered: %w", err)
	}

	test := tests[testNum-1]

	return test, nil
}

func beginTest(stp StudentTestProvider, addCorrectAnswer func(uuid.UUID)) error {
	fmt.Printf("Test:\t\t%s\n\n", stp.GetTitle())
	for n, question := range stp.GetQuestions() {

		fmt.Printf("Question %d:\t%s\n\n", n+1, question.GetText())
		a := []uuid.UUID{}
		n := 1
		for id, answer := range question.GetAnswerOptions() {
			fmt.Printf("%d) %s\n", n, answer)
			a = append(a, id)
			n++
		}
		fmt.Println()

		fmt.Print("Entry your answer: ")
		var stdAnswer int
		_, err := fmt.Scan(&stdAnswer)
		if err != nil {
			return fmt.Errorf("beginTest: invalid command value entered: %w", err)
		}

		if question.IsCorrectAnswer(a[stdAnswer]) {
			addCorrectAnswer(a[stdAnswer])
		}

		fmt.Println()
	}

	return nil
}

func addCorrectAnswer(st StudentTest) func(uuid.UUID) {
	return func(id uuid.UUID) {
		st.AddCorrectAnswer(id)
	}
}

func showResult(st StudentTest) {
	fmt.Printf("Number of correct answers: \t%d\n", st.GetCorrectAnswerCount())
	fmt.Printf("Number of wrong answers: \t%d\n", st.GetWrongAnswerCount())
}
