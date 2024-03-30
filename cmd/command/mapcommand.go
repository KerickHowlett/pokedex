package command

import (
	"fmt"

	m "github.com/KerickHowlett/pokedexcli/cmd/mapslist"
)

type MapCommand struct {
	state *m.MapsList
}

// Execute is a method of the MapCommand struct responsible for cycling through
// the pokemon world map locations list via the Pokemon API.
func (c *MapCommand) Execute() error {
	url := c.state.NextURL

	if url == nil {
		return fmt.Errorf("no more maps to fetch")
	}

	if err := c.state.FetchMapsList(*url); err != nil {
		return err
	}

	return c.state.PrintLocations()
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
func NewMapCommand(state *m.MapsList) *MapCommand {
	return &MapCommand{
		state: state,
	}
}
