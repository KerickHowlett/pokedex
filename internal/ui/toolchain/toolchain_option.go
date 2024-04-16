package toolchain

import c "command"

// ToolchainOption is a function type that represents an option for configuring
// a Toolchain.
type ToolchainOption func(*Toolchain)

// WithCommand is a function that adds a CLI command to the toolchain.
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
		commandName := command.GetName()

		if _, exists := (*t.commands)[commandName]; exists {
			panic("Command already exists in the toolchain")
		}

		(*t.commands)[commandName] = &command
	}
}
