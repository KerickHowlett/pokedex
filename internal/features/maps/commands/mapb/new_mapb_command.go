package mapb_commands

import (
	h "maps/commands/helpers"
	ms "maps/state"
)

// NewMapBCommand creates a new instance of the MapB struct.
//
// Parameters:
//   - state: A pointer to a MapsState struct instance.
//
// Returns:
//   - A new instance of the MapB struct.
//
// Example usage:
//
//	command := NewMapBCommand()
func NewMapBCommand(state ...*ms.MapsState) *MapBCommand {
	return &MapBCommand{
		state: h.InitializeStateIfNeeded(state...),
	}
}
