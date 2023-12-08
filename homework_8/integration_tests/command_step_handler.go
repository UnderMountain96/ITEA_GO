package main

import (
	"context"
	"errors"
	"strings"

	"github.com/cucumber/godog"
	"github.com/greeflas/itea_golang/cmd"
	"github.com/greeflas/itea_golang/params"
)

type CommandStepHandler struct {
	registry *cmd.Registry
}

func NewCommandStepHandler(registry *cmd.Registry) *CommandStepHandler {
	return &CommandStepHandler{registry: registry}
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
