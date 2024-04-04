package command

import (
	"fmt"

	qf "query/fetch"
	qs "query/state"

	l "github.com/KerickHowlett/pokedexcli/cmd/location"
)

type MapCommand struct {
	query *qs.QueryState[l.Location]
}

func (m *MapCommand) PrintLocationNames() error {
	if len(m.query.Results) == 0 {
		return fmt.Errorf("no maps found")
	}

	for _, location := range m.query.Results {
		fmt.Println(location.Name)
	}

	return nil
}

// Execute is a method of the MapCommand struct responsible for cycling through
// the pokemon world map locations list via the Pokemon API.
func (c *MapCommand) Execute() error {
	if c.query.NextURL == nil {
		return fmt.Errorf("no more maps to fetch")
	}

	if err := qf.QueryFetch(*c.query.NextURL, c.query); err != nil {
		return err
	}

	return c.PrintLocationNames()
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
// Parameters:
//   - state: A pointer to a QueryState struct instance.
//
// Returns:
//   - A new instance of the MapCommand struct.
//
// Example usage:
//
//	command := NewMapCommand()
func NewMapCommand(state *qs.QueryState[l.Location]) *MapCommand {
	return &MapCommand{
		query: state,
	}
}
