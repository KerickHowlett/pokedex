package maps_commands

import (
	c "command"
	"fmt"
	ms "maps/state"
	qf "query/fetch"
	"time"
)

// FetchLocations is a function that fetches locations/maps from the Pokemon API.
type FetchLocations qf.QueryFetchFunc[ms.MapsState]

// MapsCommand represents a command related to maps in the pokedexcli application.
//
// Fields:
//   - cacheTTL: The time-to-live (TTL) determines how long fetchLocations' cached responses get to exist before being discarded.
//   - description: Describes the purpose of the MapsCommand.
//   - config: A pointer to a MapsCommandsConfig struct instance.
//   - fetchLocations: A function that fetches locations/maps from the Pokemon API.
//   - listMarker: A string marker to display before each map name for nicer formatting.
//   - listTitle: A string title to display before the list of maps.
//   - name: The name of the MapsCommand.
//   - noMapsFoundErrorMessage: An error message to display when no maps are found.
//   - noMoreMapsMessage: A message to display when there are no more maps to be found.
//   - paginationDirection: Sets which direction of the maps pagination flows: "next" or "previous".
//   - state: A pointer to a MapsState struct instance.
type MapsCommand struct {
	// cacheTTL is the time-to-live (TTL) determines how long fetchLocations' cached responses get to exist before being discarded.
	cacheTTL time.Duration
	// description describes the purpose of the MapsCommand.
	description string
	// fetchLocations is a function that fetches locations/maps from the Pokemon API.
	fetchLocations FetchLocations
	// listMarker is a string marker to display before each map name for nicer formatting.
	listMarker string
	// listTitle is a string title to display before the list of maps.
	listTitle string
	// name is the name of the MapsCommand.
	name string
	// noMapsFoundErrorMessage is an error message to display when no maps are found.
	noMapsFoundErrorMessage string
	// noMoreMapsMessage is a message to display when there are no more maps to be found.
	noMoreMapsMessage string
	// paginationDirection sets which direction of the maps pagination flows: "next" or "previous".
	paginationDirection string
	// state is a pointer to a MapsState struct instance.
	state *ms.MapsState
}

// Execute executes the MapsCommand and fetches locations using the provided
// state and query.
//
// Returns:
//   - An error if the command execution fails.
//
// Example usage:
//
// m := NewMapsCommand()
// err := m.Execute()
func (m *MapsCommand) Execute() error {
	url := m.getMapsAPIEndpoint()
	if url == nil {
		fmt.Println(m.noMoreMapsMessage)
		return nil
	}

	if err := m.fetchLocations(*url, m.state, m.cacheTTL); err != nil {
		return err
	}

	if len(m.state.Locations) == 0 {
		return fmt.Errorf(m.noMapsFoundErrorMessage)
	}

	fmt.Println(m.listTitle)
	for _, location := range m.state.Locations {
		fmt.Printf("%s %s\n", m.listMarker, location.Name)
	}

	return nil
}

// GetArgs returns the arguments of the MapsCommand.
//
// This is just to appease the Command interface static type checks where needed.
// This shouldn't actually be used for anything.
func (m MapsCommand) GetArgs() []string {
	return []string{}
}

// GetDescription returns the description of the MapsCommand.
//
// Returns:
//   - A string describing the purpose of the MapsCommand.
func (m MapsCommand) GetDescription() string {
	return m.description
}

// GetName returns the name of the MapsCommand.
//
// Returns:
//   - A string containing the name of the MapsCommand.
func (m MapsCommand) GetName() string {
	return m.name
}

// PrintHelp prints the help message for the MapsCommand.
//
// Example usage:
//
//	m := NewMapsCommand()
//	m.PrintHelp()
func (m *MapsCommand) PrintHelp() {
	c.PrintHelp(m)
}

// SetArgs sets the arguments of the MapsCommand.
//
// This is just to appease the Command interface static type checks where needed.
// This shouldn't actually be used for anything.
func (m *MapsCommand) SetArgs(args []string) {}

// getMapsAPIEndpoint returns the URL of the next or previous set of maps
// from the Pokemon API.
//
// The Map Pagination Direction constants are to be treated like an enum when
// selecting the pagination direction for the command. If neither direction is
// chosen, the function will panic.
//
// Returns:
//   - A pointer to a string containing the URL of the next or previous set of maps. Nullable.
//
// Example usage:
//
// m := NewMapsCommand()
// url := m.getMapsAPIEndpoint()
func (m MapsCommand) getMapsAPIEndpoint() *string {
	switch m.paginationDirection {
	case Next:
		return m.state.NextURL
	case Previous:
		return m.state.PreviousURL
	default:
		panic("Invalid pagination direction")
	}
}
