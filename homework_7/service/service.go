package service

import (
	"fmt"

	"github.com/UnderMountain96/ITEA_GO/student_testing/model"
	"github.com/google/uuid"
)

type TestRetriever interface {
	GetTest() *model.Test
	GetTitle() string
}

type QuestionManager interface {
	GetQuestions() []*model.Question
	SetQuestions(questions []*model.Question)
}

type StudentTest interface {
	GetCorrectAnswerCount() int
	GetWrongAnswerCount() int
	AddCorrectAnswer(uuid.UUID)
}

type StudentTestProvider interface {
	TestRetriever
	QuestionManager
	StudentTest
}

func ShowAvailableTests(tr ...TestRetriever) (*model.Test, error) {
	for idx, t := range tr {
		fmt.Printf("%d) %s.\n", idx+1, t.GetTitle())
	}

	fmt.Printf("Choose test: ")
	var testNum int
	_, err := fmt.Scan(&testNum)
	if err != nil {
		return nil, fmt.Errorf("ShowAvailableTests: invalid number test value entered: %w", err)
	}

	idx := testNum - 1

	if len(tr)-1 < idx {
		return nil, fmt.Errorf("ShowAvailableTests: invalid number test value entered: index out of range [%d] with length %d", len(tr), testNum)
	}

	test := tr[idx]

	return test.GetTest(), nil
}

func BeginTest(stp StudentTestProvider) error {
	fmt.Printf("Test:\t\t%s\n\n", stp.GetTitle())
	for n, question := range stp.GetQuestions() {

		fmt.Printf("Question %d:\t%s\n\n", n+1, question.GetText())
		a := []uuid.UUID{}
		n := 1
		questions := question.GetAnswerOptions()
		for id, answer := range questions {
			fmt.Printf("%d) %s\n", n, answer)
			a = append(a, id)
			n++
		}
		fmt.Println()

		fmt.Print("Entry your answer: ")
		var stdAnswer int
		_, err := fmt.Scan(&stdAnswer)
		if err != nil {
			return fmt.Errorf("BeginTest: invalid command value entered: %w", err)
		}

		idx := stdAnswer - 1

		if len(questions)-1 < idx {
			return fmt.Errorf("BeginTest: invalid number test value entered: index out of range [%d] with length %d", len(questions), stdAnswer)
		}

		if question.IsCorrectAnswer(a[idx]) {
			stp.AddCorrectAnswer(a[idx])
		}

		fmt.Println()
	}

	return nil
}

func ShowResult(st StudentTest) {
	fmt.Printf("Number of correct answers: \t%d\n", st.GetCorrectAnswerCount())
	fmt.Printf("Number of wrong answers: \t%d\n", st.GetWrongAnswerCount())
}
