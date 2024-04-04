package command

import (
	"fmt"

	l "location"
	qf "query/fetch"
	qs "query/state"
)

type MapBCommand struct {
	query *qs.QueryState[l.Location]
}

func (m *MapBCommand) PrintLocationNames() error {
	if len(m.query.Results) == 0 {
		return fmt.Errorf("no maps found")
	}

	for _, location := range m.query.Results {
		fmt.Println(location.Name)
	}

	return nil
}

// Execute is a method of the MapBCommand struct responsible for cycling back
// through the pokemon world map locations list via the Pokemon API.
func (m *MapBCommand) Execute() error {
	if m.query.NextURL == nil {
		return fmt.Errorf("this is the start of the maps list")
	}

	if err := qf.QueryFetch(*m.query.NextURL, m.query); err != nil {
		return err
	}

	return m.PrintLocationNames()
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
func NewMapBCommand(query *qs.QueryState[l.Location]) *MapBCommand {
	return &MapBCommand{
		query: query,
	}
}
