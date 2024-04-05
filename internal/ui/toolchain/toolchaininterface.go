package toolchain

import c "command"

type ToolchainInterface interface {
	PrintHelpMessage()
	RunCommand(selection string)
	SelectCommand(commandName string) (foundCommand *c.Command, commandExists bool)
}
