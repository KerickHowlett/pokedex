package commands

import (
	c "command"
	mf "maps/fetchlocations"
	ml "maps/location"
	qs "query/state"
)

type Map struct {
	state *qs.QueryState[ml.Location]
}

func (m *Map) Execute() error {
	return mf.FetchLocations(m.state.NextURL, m.state)
}

func (m Map) GetDescription() string {
	return "Show the next 20 Pokemon world map locations."
}

func (m Map) GetName() string {
	return "map"
}

func (m *Map) PrintHelp() {
	c.PrintHelp(m)
}

// NewMap creates a new instance of the Map struct.
//
// Parameters:
//   - state: A pointer to a QueryState struct instance.
//
// Returns:
//   - A new instance of the Map struct.
//
// Example usage:
//
//	command := NewMap()
func NewMap(state *qs.QueryState[ml.Location]) *Map {
	return &Map{
		state: state,
	}
}
