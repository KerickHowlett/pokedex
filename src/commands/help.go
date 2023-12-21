package commands

import (
	"fmt"
)

func PrintHelp(commands *[]CLICommand) error {
	for _, cmd := range *commands {
		fmt.Printf("%s: %s\n", cmd.Name, cmd.Description)
	}

	return nil
}
