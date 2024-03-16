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

// SelectCommand plucks a CLI command from the toolchain based on the given
// command name.
//
// It returns the found command and a boolean indicating whether the command
// exists.
//
// Parameters:
//   - commandName: The name of the command to select.
//
// Returns:
//   - The found command.
//   - A boolean indicating whether the command exists.
//
// Example usage:
//
//	toolchain := NewToolchain(
//		WithCommand(c.NewExitCommand()),
//		WithCommand(c.NewHelpCommand()),
//	)
func (t *Toolchain) SelectCommand(commandName string) (foundCommand *c.Command, commandExists bool) {
	foundCommand, commandExists = (*t.commands)[commandName]
	return foundCommand, commandExists
}

// NewToolchain is a function that creates a new toolchain of CLI commands.
// The toolchain is used to execute commands in the REPL.
//
// Parameters:
//   - options: The options to configure the toolchain.
//
// Returns:
//   - A new instance of the Toolchain struct.
//
// Example usage:
//
//	toolchain := NewToolchain(
//		WithCommand(c.NewExitCommand()),
//		WithCommand(c.NewHelpCommand()),
//	)
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

// WithCommand is a function that adds a CLI command to the toolchain.
// The ToolchainOption function modifies the toolchain by adding the command
// to it.
//
// Parameters:
//   - command: The command to add to the toolchain.
//
// Returns:
//   - A ToolchainOption function that modifies the toolchain by adding the
//     command
//
// Example usage:
//
//	toolchain := NewToolchain(
//		WithCommand(c.NewExitCommand()),
//		WithCommand(c.NewHelpCommand()),
//	)
func WithCommand(command c.Command) ToolchainOption {
	return func(t *Toolchain) {
		commandName := command.GetName()
		(*t.commands)[commandName] = &command
	}
}
