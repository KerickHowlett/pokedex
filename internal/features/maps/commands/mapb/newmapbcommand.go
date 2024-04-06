package mapbcommands

import (
	h "maps/commands/helpers"
	ml "maps/location"
	qs "query/state"
)

// NewMapBCommand creates a new instance of the MapB struct.
//
// Parameters:
//   - state: A pointer to a QueryState struct instance.
//
// Returns:
//   - A new instance of the MapB struct.
//
// Example usage:
//
//	command := NewMapBCommand()
func NewMapBCommand(state ...*qs.QueryState[ml.Location]) *MapBCommand {
	return &MapBCommand{
		state: h.InitializeStateIfNeeded(state...),
	}
}
