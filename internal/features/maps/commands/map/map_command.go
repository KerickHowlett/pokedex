package map_commands

import (
	c "command"
	mf "maps/fetch_locations"
	ms "maps/state"
)

// MapCommand represents a command related to maps in the pokedexcli application.
type MapCommand struct {
	state *ms.MapsState
}

// Execute executes the MapCommand and fetches locations using the provided state and query.
func (m *MapCommand) Execute() error {
	return mf.FetchLocations(m.state.NextURL, m.state)
}

// GetArgs returns the arguments of the MapCommand.
func (m MapCommand) GetArgs() []string {
	return []string{}
}

// GetDescription returns the description of the MapCommand.
func (m MapCommand) GetDescription() string {
	return "Show the next 20 Pokemon world map locations."
}

// GetName returns the name of the MapCommand.
func (m MapCommand) GetName() string {
	return "map"
}

// PrintHelp prints the help message for the MapCommand.
func (m *MapCommand) PrintHelp() {
	c.PrintHelp(m)
}

// SetArgs sets the arguments of the MapCommand.
func (m *MapCommand) SetArgs(args []string) {}
