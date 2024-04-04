package command

import (
	"fmt"

	l "location"
	qf "query/fetch"
	qs "query/state"
)

type MapBCommand struct {
	state *qs.QueryState[l.Location]
}

// Execute is a method of the MapBCommand struct responsible for cycling back
// through the pokemon world map locations list via the Pokemon API.
func (m *MapBCommand) Execute() error {
	if m.state.NextURL == nil {
		return fmt.Errorf("no more maps to fetch")
	}

	if err := qf.QueryFetch(*m.state.NextURL, m.state); err != nil {
		return err
	}

	if len(m.state.Results) == 0 {
		return fmt.Errorf("no maps found")
	}

	for _, location := range m.state.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func (m MapBCommand) GetDescription() string {
	return "Show the previous 20 Pokemon world map locations."
}

func (m MapBCommand) GetName() string {
	return "mapb"
}

func (m *MapBCommand) PrintHelp() {
	printSingleCommandHelp(m)
}

// NewMapBCommand creates a new instance of the MapBCommand struct.
//
// Parameters:
//   - state: A pointer to a QueryState struct instance.
//
// Returns:
//   - A new instance of the MapBCommand struct.
//
// Example usage:
//
//	command := NewMapBCommand()
func NewMapBCommand(state *qs.QueryState[l.Location]) *MapBCommand {
	return &MapBCommand{
		state: state,
	}
}
