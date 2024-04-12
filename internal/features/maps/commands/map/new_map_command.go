package map_commands

import (
	h "maps/commands/helpers"
	ms "maps/state"
)

// NewMapCommand creates a new instance of the Map struct.
//
// Parameters:
//   - state: A pointer to a MapsState struct instance.
//
// Returns:
//   - A new instance of the Map struct.
//
// Example usage:
//
//	command := NewMapCommand()
func NewMapCommand(state ...*ms.MapsState) *MapCommand {
	return &MapCommand{
		state: h.InitializeStateIfNeeded(state...),
	}
}
