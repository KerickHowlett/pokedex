package toolchain

import "fmt"

// RunCommand executes the specified command based on the given selection.
//
// It returns an error if the selection is not valid or if there is an error
// executing the command.
//
// Parameters:
//   - selection: The name of the command to execute.
//   - args: A slice of strings representing the arguments to the command.
//
// Returns:
//   - An error if the selection is not valid or if there is an error executing
//     the command.
//
// Example usage:
//
//	toolchain := NewToolchain(
//		WithCommand(c.NewExitCommand()),
//		WithCommand(c.NewHelpCommand()),
//	)
//
//	err := toolchain.RunCommand("exit")
func (t *Toolchain) RunCommand(selection string, args []string) error {
	if selection == "" {
		return nil
	}

	if selection == "help" {
		t.PrintHelpMessage()
		return nil
	}

	if selectedCommand, commandFound := t.SelectCommand(selection); commandFound {
		(*selectedCommand).SetArgs(args)
		return (*selectedCommand).Execute()
	}

	return fmt.Errorf("[ERROR] Command '%s' is not valid", selection)
}
