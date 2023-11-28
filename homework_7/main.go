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

func makeConnectString(user, pass, host, port, db_name string) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, pass, host, port, db_name)
}
