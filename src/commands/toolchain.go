package commands

type IToolchain interface {
	PrintHelpMessage()
}

type Toolchain struct {
	commands *Commands
}

func (t *Toolchain) AddCommand(command ...Command) {
	for _, cmd := range command {
		(*t.commands)[cmd.Name] = cmd
	}
}

func (t *Toolchain) DeleteCommand(name string) {
	delete(*t.commands, name)
}

func NewToolchain() *Toolchain {
	return &Toolchain{
		commands: &Commands{},
	}
}
