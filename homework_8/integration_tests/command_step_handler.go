package main

import (
	"context"
	"errors"
	"github.com/cucumber/godog"
	"github.com/greeflas/itea_golang/cmd"
)

type CommandStepHandler struct {
	registry *cmd.Registry
}

func NewCommandStepHandler(registry *cmd.Registry) *CommandStepHandler {
	return &CommandStepHandler{registry: registry}
}

func (h *CommandStepHandler) RegisterSteps(ctx *godog.ScenarioContext) {
	ctx.Step(`^I run "([^"]*)" command$`, h.iRunCommand)
}

func (h *CommandStepHandler) iRunCommand(cmdName string) error {
	command := h.registry.FindCommand(cmdName)

	if command == nil {
		return errors.New("command not found")
	}

	return command.Run(context.Background())
}
