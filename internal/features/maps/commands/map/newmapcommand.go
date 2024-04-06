package mapcommands

import (
	h "maps/commands/helpers"
	ml "maps/location"
	qs "query/state"
)

// NewMapCommand creates a new instance of the Map struct.
//
// Parameters:
//   - state: A pointer to a QueryState struct instance.
//
// Returns:
//   - A new instance of the Map struct.
//
// Example usage:
//
//	command := NewMapCommand()
func NewMapCommand(state ...*qs.QueryState[ml.Location]) *MapCommand {
	return &MapCommand{
		state: h.InitializeStateIfNeeded(state...),
	}
}
