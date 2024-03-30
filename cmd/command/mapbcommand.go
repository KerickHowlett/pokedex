package command

import (
	"fmt"

	m "github.com/KerickHowlett/pokedexcli/cmd/mapslist"
)

type MapBCommand struct {
	state *m.MapsList
}

// Execute is a method of the MapBCommand struct responsible for cycling back
// through the pokemon world map locations list via the Pokemon API.
func (c *MapBCommand) Execute() error {
	url := c.state.PreviousURL

	if url == nil {
		return fmt.Errorf("this is the start of the maps list")
	}

	if err := c.state.FetchMapsList(*url); err != nil {
		return err
	}

	return c.state.PrintLocations()
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
func NewMapBCommand(state *m.MapsList) *MapBCommand {
	return &MapBCommand{
		state: state,
	}
}
