package toolchain

import (
	c "github.com/KerickHowlett/pokedexcli/src/toolchain/command"
)

type IToolchain interface {
	PrintHelpMessage()
	RunCommand(selection string)
	SelectCommand(commandName string) (foundCommand *c.Command, commandExists bool)
}

type Toolchain struct {
	commands *c.Commands
}

type ToolchainOption func(*Toolchain)

func (t *Toolchain) SelectCommand(commandName string) (foundCommand *c.Command, commandExists bool) {
	foundCommand, commandExists = (*t.commands)[commandName]
	return foundCommand, commandExists
}

func NewToolchain(options ...ToolchainOption) *Toolchain {
	toolchain := &Toolchain{
		commands: &c.Commands{},
	}

	for _, option := range options {
		option(toolchain)
	}

	if len(*toolchain.commands) == 0 {
		panic("[ERROR] Toolchain must have at least one command.")
	}

	return toolchain
}

func WithCommand(command c.Command) ToolchainOption {
	return func(t *Toolchain) {
		commandName := command.GetName()
		(*t.commands)[commandName] = &command
	}
}
