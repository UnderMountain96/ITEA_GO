package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/UnderMountain96/ITEA_GO/student_testing/dotenv"
	"github.com/UnderMountain96/ITEA_GO/student_testing/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type StudentTest interface {
	GetCorrectAnswerCount() int
	GetWrongAnswerCount() int
	AddCorrectAnswer(id uuid.UUID)
}

type StudentTestProvider interface {
	GetTitle() string
	GetQuestions() []model.Question
}

func main() {
	loadEnv()

	ctx := context.Background()

	connStr := makeConnectString(
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	fmt.Println(connStr)
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	rows, err := conn.Query(ctx, "SELECT id, title FROM test")
	if err != nil {
		fmt.Printf("Query error: %s\n", err)
	}

	availableTests := make([]model.Test, 0)

	for rows.Next() {
		var id, title string
		if err := rows.Scan(&id, &title); err != nil {
			fmt.Printf("Scan error: %s\n", err)
		}
		t := model.NewTest(uuid.MustParse(id), title, []model.Question{})

		availableTests = append(availableTests, t)
	}

	showAvailableTests(availableTests)

	// beginTest(test, addCorrectAnswer(test))

	// showResult(test)
}

func showAvailableTests(tests []model.Test) (uuid.UUID, error) {
	for idx, t := range tests {
		fmt.Printf("%d) %s.\n", idx+1, t.GetTitle())
	}

	fmt.Printf("Choose test: ")
	var testNum int
	_, err := fmt.Scan(&testNum)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("showAvailableTests: invalid number test value entered: %w", err)
	}

	selected := tests[testNum-1]

	fmt.Printf("%s\n", selected.GetTitle())

	return selected.GetID(), nil
}

func loadEnv() {
	const envFilePath = "./.env"

	err := dotenv.LoadEnv(envFilePath)

	if errors.Is(err, os.ErrNotExist) {
		fmt.Printf("Oops, looks like file %q does not exist. You need to create it.\n", envFilePath)

		return
	}

	if err != nil {
		fmt.Printf("Fatal error: %s\n", err)

		return
	}
}

func makeConnectString(user, pass, host, port, db_name string) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, pass, host, port, db_name)
}

func beginTest(stp StudentTestProvider, addCorrectAnswer func(int)) {
	fmt.Printf("Test:\t\t%s\n\n", stp.GetTitle())
	// for n, question := range stp.GetQuestions() {
	// 	questionNumner := n + 1

	// 	fmt.Printf("Question %d:\t%s\n\n", questionNumner, question.Text)
	// 	for idx, answer := range question.AnswerOptions {
	// 		fmt.Printf("%d) %s\n", idx, answer)
	// 	}
	// 	fmt.Println()

	// 	fmt.Print("Entry your answer: ")
	// 	var stdAnswer int
	// 	fmt.Scan(&stdAnswer)

	// 	if question.IsCorrectAnswer(stdAnswer) {
	// 		addCorrectAnswer(questionNumner)
	// 	}

	// 	fmt.Println()
	// }
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
