package explore

import (
	c "command"
	fe "explore/fetch_encounters"
	la "explore/location_area"
	"fmt"
)

// ExploreCommand represents a command related to maps in the pokedexcli application.
//
// Fields:
//   - args: The arguments provided to the command.
//   - state: The state of the location area to be explored.
type ExploreCommand struct {
	// The arguments provided to the command.
	args []string
	// The URL of the API endpoint to fetch the location area data.
	apiEndpoint string
	// The state of the location area to be explored.
	state *la.LocationArea
}

// Execute executes the MapCommand and fetches locations using the provided state and query.
func (m *ExploreCommand) Execute() error {
	if m.apiEndpoint == "" {
		return fmt.Errorf("an API endpoint is required to fetch location areas")
	}

	if len(m.args) == 0 {
		return fmt.Errorf("a location area name is required")
	}

	locationAreaEndpoint := m.apiEndpoint + "/" + m.args[0]
	return fe.FetchEncounters(locationAreaEndpoint, m.state)
}

// GetArgs returns the arguments of the ExploreCommand.
func (m ExploreCommand) GetArgs() []string {
	return m.args
}

// GetDescription returns the description of the MapCommand.
func (m ExploreCommand) GetDescription() string {
	return "Show the Pokemon that can be encountered in the location area entered. Usage: explore <location-area-name>"
}

// GetName returns the name of the MapCommand.
func (m ExploreCommand) GetName() string {
	return "explore"
}

// PrintHelp prints the help message for the MapCommand.
func (m *ExploreCommand) PrintHelp() {
	c.PrintHelp(m)
}

// SetArgs sets the arguments of the ExploreCommand.
func (m *ExploreCommand) SetArgs(args []string) {
	m.args = args
}
