package command

import "fmt"

type MapBCommand struct{}

// Execute is a method of the MapBCommand struct responsible for cycling back
// through the pokemon world map locations list via the Pokemon API.
func (c *MapBCommand) Execute() error {
	fmt.Println("TODO: Implement 'map' command")
	return nil
}

func (c *MapBCommand) GetDescription() string {
	return "Show the previous 20 Pokemon world map locations."
}

func (c *MapBCommand) GetName() string {
	return "mapb"
}

func (c *MapBCommand) PrintHelp() {
	printSingleCommandHelp(c)
}

// NewMapBCommand creates a new instance of the MapBCommand struct.
//
// Returns:
//   - A new instance of the MapBCommand struct.
//
// Example usage:
//
//	command := NewMapBCommand()
func NewMapBCommand() *MapBCommand {
	return &MapBCommand{}
}
