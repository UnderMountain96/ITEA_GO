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

	connStr := "postgres://postgres:pass@192.168.230.128:5432/lessons"
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	articleRepository := repository.NewArticleRepository(conn)
	createArticleCommand := cmd.NewCreateArticleCommand(articleRepository)
	updateArticleCommand := cmd.NewUpdateArticleCommand(articleRepository)
	commandRegistry := cmd.NewRegistry(createArticleCommand, updateArticleCommand)

	suite := godog.TestSuite{
		Name: "Articles agency",
		Options: &godog.Options{
			Format: "pretty",
			Paths: []string{
				"features",
			},
		},
		ScenarioInitializer: func(ctx *godog.ScenarioContext) {
			ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
				_, err := conn.Exec(
					context.Background(),
					"INSERT INTO articles (id, title, body) VALUES ($1, $2, $3)",
					"a462db9b-b7ae-434c-87af-943d080d5c00",
					"for update",
					"some body",
				)

				return ctx, err
			})

			ctx.After(func(ctx context.Context, sc *godog.Scenario, err error) (context.Context, error) {
				_, err = conn.Exec(
					context.Background(),
					"DELETE FROM articles",
				)

				return ctx, err
			})

			commandStepHandler := NewCommandStepHandler(commandRegistry)
			commandStepHandler.RegisterSteps(ctx)

			pgxStepHandler := NewPgxStepHandler(conn)
			pgxStepHandler.RegisterSteps(ctx)
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature test")
	}
}
