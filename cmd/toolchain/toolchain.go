package toolchain

import c "github.com/KerickHowlett/pokedexcli/cmd/command"

type IToolchain interface {
	PrintHelpMessage()
	RunCommand(selection string)
	SelectCommand(commandName string) (foundCommand *c.Command, commandExists bool)
}

type Toolchain struct {
	commands *c.Commands
}

type ToolchainOption func(*Toolchain)

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
