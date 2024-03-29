package command

import "fmt"

type MapCommand struct{}

// Execute is a method of the MapCommand struct responsible for cycling through
// the pokemon world map locations list via the Pokemon API.
func (c *MapCommand) Execute() error {
	fmt.Println("TODO: Implement 'map' command")
	return nil
}

func (c *MapCommand) GetDescription() string {
	return "Show the next 20 Pokemon world map locations."
}

func (c *MapCommand) GetName() string {
	return "map"
}

func (c *MapCommand) PrintHelp() {
	printSingleCommandHelp(c)
}

// NewMapCommand creates a new instance of the MapCommand struct.
//
// Returns:
//   - A new instance of the MapCommand struct.
//
// Example usage:
//
//	command := NewMapCommand()
func NewMapCommand() *MapCommand {
	return &MapCommand{}
}
