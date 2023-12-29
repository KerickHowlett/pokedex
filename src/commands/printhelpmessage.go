package commands

import "fmt"

// Prints the help message for the Pokedex CLI application.
//
// It displays the welcome message and usage instructions, and it prints the
// help information for each command.
func PrintHelpMessage() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println() // Empty Line

	fmt.Println("Pokedex Commands")

	commands := GetCommands()
	printCommandsHelp(commands)

	return nil
}

func printCommandsHelp(commands Commands) {
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.Name, command.Description)
	}
}
