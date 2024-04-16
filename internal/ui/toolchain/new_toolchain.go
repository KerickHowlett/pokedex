package toolchain

import (
	c "command"
	hmp "toochain/help_message_prints"
)

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
		prints:   hmp.NewHelpMessagePrints(),
	}

	for _, option := range options {
		option(toolchain)
	}

	if len(*toolchain.commands) == 0 {
		panic("toolchain must have at least one command")
	}

	return toolchain
}
