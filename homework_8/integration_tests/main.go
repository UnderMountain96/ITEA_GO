package main

import (
	"context"
	"flag"
	"os"

	"github.com/greeflas/itea_golang/cmd"
	"github.com/greeflas/itea_golang/params"
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
	updateArticleCommand := cmd.NewUpdateArticleCommand(articleRepository)
	deleteArticleCommand := cmd.NewDeleteArticleCommand(articleRepository)
	commandRegistry := cmd.NewRegistry(createArticleCommand, getAllArticleCommand, updateArticleCommand, deleteArticleCommand)

	cmdName := os.Args[1]
	command := commandRegistry.FindCommand(cmdName)
	if command == nil {
		panic("command not found")
	}

	cmdParams := flag.NewFlagSet("params", flag.ExitOnError)
	var params params.MapParams
	cmdParams.Var(&params, "p", "Map flag usage: -p id=id -p body=newBody -p title=newTitle")

	cmdParams.Parse(os.Args[2:])

	if err := command.Run(ctx, params); err != nil {
		panic(err)
	}
}
