package command

type Command interface {
	Execute() error
	GetDescription() string
	GetName() string
	PrintHelp()
}

type Commands = map[string]*Command
