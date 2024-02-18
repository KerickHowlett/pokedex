package command

import "fmt"

func printSingleCommandHelp(command Command) {
	name := command.GetName()
	description := command.GetDescription()

	fmt.Printf("%s: %s\n", name, description)
}
