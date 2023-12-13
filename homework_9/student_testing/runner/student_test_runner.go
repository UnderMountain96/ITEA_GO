package runner

import (
	"fmt"

	"github.com/UnderMountain96/ITEA_GO/student_testing_json/model"
)

type TestProvider interface {
	GetTests() ([]*model.Test, error)
}

type StudentTestRunner struct {
	testsProvider []TestProvider
}

func NewStudentTestRunner(testsProvider ...TestProvider) *StudentTestRunner {
	return &StudentTestRunner{testsProvider: testsProvider}
}

func (r *StudentTestRunner) Run() error {
	allTests := map[int]*model.Test{}
	n := 1
	for _, tp := range r.testsProvider {
		tests, err := tp.GetTests()
		if err != nil {
			return fmt.Errorf("run: cannot get tests from provider: %w", err)
		}
		for _, answer := range tests {
			allTests[n] = answer
			n++
		}
	}

	r.printTests(allTests)

	test, err := r.askTest(allTests)
	if err != nil {
		return err
	}

	correctAnswersCount, incorrectAnswersCount, err := r.askQuestions(test.GetQuestions())
	if err != nil {
		return err
	}

	r.printResult(correctAnswersCount, incorrectAnswersCount)

	return nil
}

func (r *StudentTestRunner) askTest(tests map[int]*model.Test) (*model.Test, error) {
	for {
		fmt.Print("\nYour variant > ")

		var numTest int
		if _, err := fmt.Scan(&numTest); err != nil {
			return nil, fmt.Errorf("askTest: scan error: %w", err)
		}

		answer, ok := tests[numTest]
		if ok {
			return answer, nil
		}

		fmt.Println("There is no such variant of number for this tests. Try again.")
	}
}

func (r *StudentTestRunner) askQuestions(questions []*model.Question) (correctAnswersCount int, incorrectAnswersCount int, err error) {
	for _, question := range questions {
		answerOptions := r.createAnswerOptions(question)

		r.printQuestionWithAnswerOptions(question, answerOptions)

		studentsAnswer, err := r.askAnswer(answerOptions)
		if err != nil {
			return 0, 0, fmt.Errorf("askQuestions: cannot get students answer: %w", err)
		}

		if question.IsCorrectAnswer(studentsAnswer.ID) {
			correctAnswersCount++
		} else {
			incorrectAnswersCount++
		}
	}

	return
}

func (r *StudentTestRunner) printResult(correctAnswersCount, incorrectAnswersCount int) {
	fmt.Printf("\nWell done!\n")
	fmt.Printf("Number of correct answers is: %d\n", correctAnswersCount)
	fmt.Printf("Number of incorrect answers is: %d\n", incorrectAnswersCount)
}

func (r *StudentTestRunner) createAnswerOptions(question *model.Question) map[string]*model.Answer {
	var answerOptions = make(map[string]*model.Answer)

	for i, answer := range question.GetAnswerOptions() {
		variantIndex := string(rune('a' + i))
		answerOptions[variantIndex] = answer
	}

	return answerOptions
}

func (r *StudentTestRunner) printTests(tests map[int]*model.Test) {
	for option, test := range tests {
		fmt.Printf("%d) %s\n", option, test.GetTitle())
	}
}

func (r *StudentTestRunner) printQuestionWithAnswerOptions(question *model.Question, options map[string]*model.Answer) {
	fmt.Printf("\n%s\n", question.GetText())

	for option, answer := range options {
		fmt.Printf("%s) %s\n", option, answer.Text)
	}
}

func (r *StudentTestRunner) askAnswer(answerOptions map[string]*model.Answer) (*model.Answer, error) {
	for {
		fmt.Print("\nYour variant > ")

		var studentsAnswer string
		if _, err := fmt.Scan(&studentsAnswer); err != nil {
			return nil, fmt.Errorf("askAnswer: scan error: %w", err)
		}

		answer, ok := answerOptions[studentsAnswer]
		if ok {
			return answer, nil
		}

		fmt.Println("There is no such variant of answer for this question. Try again.")
	}
}
