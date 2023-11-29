package main

import (
	"context"
	"testing"

	"github.com/cucumber/godog"
	"github.com/greeflas/itea_golang/cmd"
	"github.com/greeflas/itea_golang/repository"
	"github.com/jackc/pgx/v5"
)

func TestFeatures(t *testing.T) {
	ctx := context.Background()

	connStr := "postgres://postgres:pass@localhost:5432/lessons"
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	articleRepository := repository.NewArticleRepository(conn)
	createArticleCommand := cmd.NewCreateArticleCommand(articleRepository)
	getAllArticleCommand := cmd.NewGetAllArticleCommand(articleRepository)
	updateArticleCommand := cmd.NewUpdateArticleCommand(articleRepository)
	commandRegistry := cmd.NewRegistry(createArticleCommand, getAllArticleCommand, updateArticleCommand)

	suite := godog.TestSuite{
		Name: "Articles agency",
		Options: &godog.Options{
			Format: "pretty",
			Paths: []string{
				"features",
			},
		},
		ScenarioInitializer: func(context *godog.ScenarioContext) {
			commandStepHandler := NewCommandStepHandler(commandRegistry)
			commandStepHandler.RegisterSteps(context)

			pgxStepHandler := NewPgxStepHandler(conn)
			pgxStepHandler.RegisterSteps(context)
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature test")
	}
}
