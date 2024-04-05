package commands

import (
	c "command"
	mf "maps/fetchlocations"
	ml "maps/location"
	qf "query/fetch"
	qs "query/state"
)

type MapB struct {
	state *qs.QueryState[ml.Location]
}

func (m *MapB) Execute() error {
	return mf.FetchLocations(m.state.PreviousURL, m.state, qf.QueryFetch[ml.Location])
}

func (m MapB) GetDescription() string {
	return "Show the previous 20 Pokemon world map locations."
}

func (m MapB) GetName() string {
	return "mapb"
}

func (m *MapB) PrintHelp() {
	c.PrintHelp(m)
}

// NewMapB creates a new instance of the MapB struct.
//
// Parameters:
//   - state: A pointer to a QueryState struct instance.
//
// Returns:
//   - A new instance of the MapB struct.
//
// Example usage:
//
//	command := NewMapB()
func NewMapB(state *qs.QueryState[ml.Location]) *MapB {
	return &MapB{
		state: state,
	}
}
