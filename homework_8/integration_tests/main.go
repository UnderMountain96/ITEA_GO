package main

import (
	"context"
	"os"

	"github.com/greeflas/itea_golang/cmd"
	"github.com/greeflas/itea_golang/repository"
	"github.com/jackc/pgx/v5"
)

func main() {
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
	commandRegistry := cmd.NewRegistry(createArticleCommand, getAllArticleCommand)

	cmdName := os.Args[1]
	command := commandRegistry.FindCommand(cmdName)
	if command == nil {
		panic("command not found")
	}

	if err := command.Run(ctx); err != nil {
		panic(err)
	}
}
