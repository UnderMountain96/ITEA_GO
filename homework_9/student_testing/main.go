package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"io"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/UnderMountain96/ITEA_GO/student_testing_json/dotenv"
	"github.com/UnderMountain96/ITEA_GO/student_testing_json/jsonstore"
	"github.com/UnderMountain96/ITEA_GO/student_testing_json/provider"
	"github.com/UnderMountain96/ITEA_GO/student_testing_json/runner"
	"github.com/UnderMountain96/ITEA_GO/student_testing_json/store"
)

func main() {
	const envFilePath = "./.env"
	if err := loadEnv(envFilePath); err != nil {
		panic(err)
	}

	db, err := newDB()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	conn, err := db.Conn(ctx)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	testsData, err := readFile("./data/tests.json")
	if err != nil {
		panic(err)
	}

	questionsData, err := readFile("./data/questions.json")
	if err != nil {
		panic(err)
	}

	testRepository := store.NewTestRepository(conn)
	questionRepository := store.NewQuestionRepository(conn)

	jsonTestRepository := jsonstore.NewTestRepository(testsData)
	jsonQuestionRepository := jsonstore.NewQuestionRepository(questionsData)

	jsonPrivider := provider.NewJsonTestsProvider(jsonTestRepository, jsonQuestionRepository)
	databasePrivider := provider.NewDatabaseTestsProvider(testRepository, questionRepository)

	studentTestRunner := runner.NewStudentTestRunner(jsonPrivider, databasePrivider)

	if err := studentTestRunner.Run(); err != nil {
		panic(err)
	}

	// availableTests, err := testRepository.GetAll(ctx)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// availableJsonTests, err := jsonTestRepository.GetAll()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// testRetriever := []service.TestRetriever{}

	// for _, t := range availableTests {
	// 	testRetriever = append(testRetriever, t)
	// }
	// for _, t := range availableJsonTests {
	// 	testRetriever = append(testRetriever, t)
	// }

	// test, err := service.ShowAvailableTests(testRetriever...)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// questions, err := questionRepository.Get(ctx, test.GetID())
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// questions1, err := jsonQuestionRepository.Get(test.GetID())
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// _ = questions1
	// test.SetQuestions(questions)

	// if err := service.BeginTest(test); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// service.ShowResult(test)
}

func loadEnv(envFilePath string) error {
	err := dotenv.LoadEnv(envFilePath)

	if errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("oops, looks like file %q does not exist. You need to create it", envFilePath)

	}

	if err != nil {
		return fmt.Errorf("fatal error: %s", err)
	}

	return nil
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

func readFile(path string) ([]byte, error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()

	return io.ReadAll(jsonFile)
}
