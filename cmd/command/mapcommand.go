package command

import (
	c "command"
	mf "maps/fetchlocations"
	ml "maps/location"
	qs "query/state"
)

type MapCommand struct {
	state *qs.QueryState[ml.Location]
}

func (m *MapCommand) Execute() error {
	return mf.FetchLocations(m.state.NextURL, m.state)
}

func (m MapCommand) GetDescription() string {
	return "Show the next 20 Pokemon world map locations."
}

func (m MapCommand) GetName() string {
	return "map"
}

func (m *MapCommand) PrintHelp() {
	c.PrintHelp(m)
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
func NewMapCommand(state *qs.QueryState[ml.Location]) *MapCommand {
	return &MapCommand{
		state: state,
	}
}
