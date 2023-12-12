package main

import (
	"context"
	"errors"
	"strings"

	"github.com/UnderMountain96/ITEA_GO/cmd"
	"github.com/UnderMountain96/ITEA_GO/params"
	"github.com/cucumber/godog"
	"github.com/jackc/pgx/v5"
)

type CommandStepHandler struct {
	registry *cmd.Registry
	conn     *pgx.Conn
}

func NewCommandStepHandler(registry *cmd.Registry, conn *pgx.Conn) *CommandStepHandler {
	return &CommandStepHandler{registry: registry, conn: conn}
}

func (h *CommandStepHandler) RegisterSteps(ctx *godog.ScenarioContext) {
	ctx.Step(`^I run "([^"]*)" command$`, h.iRunCommand)
	ctx.Step(`^I run "([^"]*)" command with params "([^"]*)"$`, h.iRunCommandWithParams)
}

func (h *CommandStepHandler) iRunCommand(cmdName string) error {
	command := h.registry.FindCommand(cmdName)

	if command == nil {
		return errors.New("command not found")
	}

	return command.Run(context.Background(), map[string]string{})
}

func (h *CommandStepHandler) iRunCommandWithParams(cmdName string, flags string) error {
	_, err := h.conn.Exec(
		context.Background(),
		"INSERT INTO articles (id, title, body) VALUES ($1, $2, $3)",
		"a462db9b-b7ae-434c-87af-943d080d5c00",
		"for update",
		"some body",
	)
	if err != nil {
		panic(err)
	}
	command := h.registry.FindCommand(cmdName)

	if command == nil {
		return errors.New("command not found")
	}

	var p params.MapParams
	flagsArr := strings.Split(flags, ",")
	for _, f := range flagsArr {
		if err := p.Set(f); err != nil {
			return err
		}
	}

	return command.Run(context.Background(), p)
}
