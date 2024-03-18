package toolchain

import c "github.com/KerickHowlett/pokedexcli/internal/command"

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
