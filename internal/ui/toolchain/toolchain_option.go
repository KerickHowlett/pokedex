package toolchain

import (
	c "command"
	hmp "toochain/help_message_prints"
)

// ToolchainOption is a function type that represents an option for configuring
// a Toolchain.
type ToolchainOption func(*Toolchain)

// WithCommand is a function that adds a CLI command to the toolchain.
//
// The ToolchainOption function modifies the toolchain by adding the command
// to it.
//
// Parameters:
//   - command: The command to add to the toolchain.
//
// Returns:
//   - A ToolchainOption function that modifies the toolchain by adding the
//     command to it.
//   - If the command already exists in the toolchain, a panic will occur.
//
// Example usage:
//
//	toolchain := NewToolchain(
//		WithCommand(c.NewExitCommand()),
//		WithCommand(c.NewHelpCommand()),
//	)
func WithCommand(command c.Command) ToolchainOption {
	return func(t *Toolchain) {
		name := command.GetName()
		if _, exists := (*t.commands)[name]; exists {
			panic("command already exists in the toolchain")
		}
		(*t.commands)[name] = &command
	}
}

// WithHelpMessagePrints is a function that sets the help message prints for
// the toolchain.
//
// The ToolchainOption function modifies the toolchain by setting the help
// message prints to the given prints.
//
// Parameters:
//   - prints: The help message prints to set for the toolchain.
//
// Returns:
//   - A ToolchainOption function that modifies the toolchain by setting the
//     help message prints to the given prints.
//
// Example usage:
//
//	toolchain := NewToolchain(
//		WithHelpMessagePrints(hmp.NewHelpMessagePrints()),
//	)
func WithHelpMessagePrints(prints *hmp.HelpMessagePrints) ToolchainOption {
	return func(t *Toolchain) {
		t.prints = prints
	}
}
