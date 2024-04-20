package map_command

import (
	"fmt"
	qec "query_fetch/query_cache/cache_eviction_config"

	c "command"
	pd "map_command/pagination_direction"
	ms "map_command/state"
	qf "query_fetch"
)

// FetchLocations is a function that fetches locations/maps from the Pokemon API.
type FetchLocations qf.QueryFetchFunc[ms.MapsState]

// MapCommand represents a command related to maps in the pokedexcli application.
//
// Fields:
//   - description: Describes the purpose of the MapCommand.
//   - ec: The eviction configuration for the query cache.
//   - fetchLocations: A function that fetches locations/maps from the Pokemon API.
//   - listMarker: A string marker to display before each map name for nicer formatting.
//   - listTitle: A string title to display before the list of maps.
//   - name: The name of the MapCommand.
//   - noMapsFoundErrorMessage: An error message to display when no maps are found.
//   - noMoreMapsMessage: A message to display when there are no more maps to be found.
//   - paginationDirection: Sets which direction of the maps pagination flows: "next" or "previous".
//   - state: A pointer to a MapsState struct instance.
//
// Methods:
//   - Execute: Executes the MapCommand and fetches locations using the provided state and query.
//   - GetArgs: Returns the arguments of the MapCommand.
//   - GetDescription: Returns the description of the MapCommand.
//   - GetName: Returns the name of the MapCommand.
//   - PrintHelp: Prints the help message for the MapCommand.
//   - SetArgs: Sets the arguments of the MapCommand.
//   - getMapsAPIEndpoint: Returns the URL of the next or previous set of maps from the Pokemon API.
type MapCommand struct {
	// description describes the purpose of the MapCommand.
	description string
	// ec is the Eviction Configuration for the query cache.
	ec *qec.QueryEvictionConfig
	// fetchLocations is a function that fetches locations/maps from the Pokemon API.
	fetchLocations FetchLocations
	// listMarker is a string marker to display before each map name for nicer formatting.
	listMarker string
	// listTitle is a string title to display before the list of maps.
	listTitle string
	// name is the name of the MapCommand.
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

// Execute executes the MapCommand and fetches locations using the provided
// state and query.
//
// Returns:
//   - An error if the command execution fails.
//
// Example usage:
//
// m := NewMapCommand()
// err := m.Execute()
func (m *MapCommand) Execute() error {
	url := m.getMapsAPIEndpoint()
	if url == nil {
		fmt.Println(m.noMoreMapsMessage)
		return nil
	}

	state, err := m.fetchLocations(*url, m.ec)
	if err != nil {
		return err
	}

	m.state = state // Save the updated state.
	if len(m.state.Locations) == 0 {
		return fmt.Errorf(m.noMapsFoundErrorMessage)
	}

	fmt.Println(m.listTitle)
	for _, location := range m.state.Locations {
		fmt.Printf("%s %s\n", m.listMarker, location.Name)
	}

	return nil
}

// GetArgs returns the arguments of the MapCommand.
//
// This is just to appease the Command interface static type checks where needed.
// This shouldn't actually be used for anything.
func (m MapCommand) GetArgs() []string {
	return []string{}
}

// GetDescription returns the description of the MapCommand.
//
// Returns:
//   - A string describing the purpose of the MapCommand.
//
// Example usage:
//
// m := NewMapCommand()
// description := m.GetDescription()
func (m MapCommand) GetDescription() string {
	return m.description
}

// GetName returns the name of the MapCommand.
//
// Returns:
//   - A string containing the name of the MapCommand.
//
// Example usage:
//
// m := NewMapCommand()
// name := m.GetName()
func (m MapCommand) GetName() string {
	return m.name
}

// PrintHelp prints the help message for the MapCommand.
//
// Example usage:
//
//	m := NewMapCommand()
//	m.PrintHelp()
func (m *MapCommand) PrintHelp() {
	c.PrintHelp(m)
}

// SetArgs sets the arguments of the MapCommand.
//
// This is just to appease the Command interface static type checks where needed.
// This shouldn't actually be used for anything.
func (m *MapCommand) SetArgs(args []string) {}

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
// m := NewMapCommand()
// url := m.getMapsAPIEndpoint()
func (m MapCommand) getMapsAPIEndpoint() *string {
	switch m.paginationDirection {
	case pd.Next:
		return m.state.NextURL
	case pd.Previous:
		return m.state.PreviousURL
	default:
		panic("invalid pagination direction")
	}
}
