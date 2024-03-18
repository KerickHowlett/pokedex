package toolchain

import c "github.com/KerickHowlett/pokedexcli/internal/command"

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
