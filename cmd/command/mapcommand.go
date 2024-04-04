package command

import (
	"fmt"

	l "location"
	qf "query/fetch"
	qs "query/state"
)

type MapCommand struct {
	state *qs.QueryState[l.Location]
}

// Execute is a method of the MapCommand struct responsible for cycling through
// the pokemon world map locations list via the Pokemon API.
func (m *MapCommand) Execute() error {
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

func (m MapCommand) GetDescription() string {
	return "Show the next 20 Pokemon world map locations."
}

func (m MapCommand) GetName() string {
	return "map"
}

func (m *MapCommand) PrintHelp() {
	printSingleCommandHelp(m)
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
		state: state,
	}
}
