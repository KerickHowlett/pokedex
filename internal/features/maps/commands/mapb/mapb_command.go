package mapb_commands

import (
	c "command"
	mf "maps/fetch_locations"
	ms "maps/state"
)

// MapBCommand represents a command for Map B feature.
type MapBCommand struct {
	state *ms.MapsState
}

// Execute fetches locations using the previous URL, state, and query fetch.
func (m *MapBCommand) Execute() error {
	return mf.FetchLocations(m.state.PreviousURL, m.state)
}

// GetDescription returns the description of the MapBCommand.
func (m MapBCommand) GetDescription() string {
	return "Show the previous 20 Pokemon world map locations."
}

// GetName returns the name of the MapBCommand.
func (m MapBCommand) GetName() string {
	return "mapb"
}

// PrintHelp prints the help message for the MapBCommand.
func (m *MapBCommand) PrintHelp() {
	c.PrintHelp(m)
}
