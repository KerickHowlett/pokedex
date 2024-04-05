package command

import (
	c "command"
	mf "maps/fetchlocations"
	ml "maps/location"
	qs "query/state"
)

type MapBCommand struct {
	state *qs.QueryState[ml.Location]
}

func (m *MapBCommand) Execute() error {
	return mf.FetchLocations(m.state.PreviousURL, m.state)
}

func (m MapBCommand) GetDescription() string {
	return "Show the previous 20 Pokemon world map locations."
}

func (m MapBCommand) GetName() string {
	return "mapb"
}

func (m *MapBCommand) PrintHelp() {
	c.PrintHelp(m)
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
func NewMapBCommand(state *qs.QueryState[ml.Location]) *MapBCommand {
	return &MapBCommand{
		state: state,
	}
}
