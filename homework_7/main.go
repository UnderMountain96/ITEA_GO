package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/UnderMountain96/ITEA_GO/student_testing/dotenv"
	"github.com/UnderMountain96/ITEA_GO/student_testing/model"
	"github.com/UnderMountain96/ITEA_GO/student_testing/store/sqlstore"
	"github.com/google/uuid"
)

type StudentTest interface {
	GetCorrectAnswerCount() int
	GetWrongAnswerCount() int
	AddCorrectAnswer(uuid.UUID)
}

type StudentTestProvider interface {
	GetTitle() string
	GetQuestions() []*model.Question
	SetQuestions([]*model.Question)
}

func main() {
	const envFilePath = "./.env"
	loadEnv(envFilePath)

	db, err := newDB()
	if err != nil {
		fmt.Println(err)
		return
	}

	dbStore := sqlstore.NewStore(db)

	ctx := context.Background()

	availableTests, err := dbStore.Test().GetAll(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}

	test, err := showAvailableTests(availableTests)
	if err != nil {
		fmt.Println(err)
		return
	}

	questions, err := dbStore.Test().GetQuestions(ctx, test.GetID())
	if err != nil {
		fmt.Println(err)
		return
	}

	test.SetQuestions(questions)

	if err := beginTest(test, addCorrectAnswer(test)); err != nil {
		fmt.Println(err)
		return
	}

	showResult(test)
}

func newDB() (*sql.DB, error) {
	connStr := makeConnectString(
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		return nil, fmt.Errorf("newDB: failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("newDB: failed to ping database: %w", err)
	}

	return db, nil
}

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

	fmt.Printf("%s\n", test.GetTitle())

	return test, nil
}

func loadEnv(envFilePath string) {
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
