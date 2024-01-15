package commands

type IToolchain interface {
	PrintHelpMessage()
}

type Toolchain struct {
	commands *Commands
}
