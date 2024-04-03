package toolchain

import c "github.com/KerickHowlett/pokedexcli/cmd/command"

type ToolchainInterface interface {
	PrintHelpMessage()
	RunCommand(selection string)
	SelectCommand(commandName string) (foundCommand *c.Command, commandExists bool)
}
