package cmd

import "context"

type Cmd interface {
	Name() string
	Run(ctx context.Context) error
}

type Registry struct {
	commands []Cmd
}

func NewRegistry(commands ...Cmd) *Registry {
	return &Registry{commands: commands}
}

func (r *Registry) FindCommand(name string) Cmd {
	for _, command := range r.commands {
		if command.Name() == name {
			return command
		}
	}

	return nil
}
