package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/UnderMountain96/ITEA_GO/cmd"
	"github.com/UnderMountain96/ITEA_GO/repository"
	"github.com/cucumber/godog"
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
				_, err = conn.Exec(
					context.Background(),
					"DELETE FROM articles",
				)

				fmt.Println("CLEAR TABLE articles")

				return ctx, err
			})

			commandStepHandler := NewCommandStepHandler(commandRegistry, conn)
			commandStepHandler.RegisterSteps(ctx)

			pgxStepHandler := NewPgxStepHandler(conn)
			pgxStepHandler.RegisterSteps(ctx)
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature test")
	}
}
