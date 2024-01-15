package commands

import "fmt"

type Command struct {
	Name        string
	Description string
	Execute     func() error
}

type ICommand interface {
	Execute() error
	PrintHelp()
}

func (c Command) PrintHelp() {
	fmt.Printf("%s: %s\n", c.Name, c.Description)
}
