package toolchain

import c "command"

// ToolchainInterface represents an interface for interacting with the toolchain.
//
// Methods:
//   - PrintHelpMessage: Prints the help message for the toolchain.
//   - RunCommand: Runs the specified command in the toolchain.
//   - SelectCommand: Selects the command with the specified name from the toolchain.
type ToolchainInterface interface {
	PrintHelpMessage()
	RunCommand(selection string)
	SelectCommand(commandName string) (foundCommand *c.Command, commandExists bool)
}
