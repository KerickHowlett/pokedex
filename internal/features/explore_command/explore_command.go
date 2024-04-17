package explore_command

import (
	"fmt"
	"time"

	c "command"
	la "location_area"
	qf "query_fetch"
)

// FetchEncounters is a type alias for the QueryFetchFunc type that fetches locations/maps from the Pokemon API.
type FetchEncounters qf.QueryFetchFunc[la.LocationArea]

// ExploreCommand represents a command related to maps in the pokedexcli application.
//
// Fields:
//   - args: The arguments provided to the command.
//   - apiEndpoint: The URL of the API endpoint to fetch the location area data.
//   - cacheTTL: The time-to-live (TTL) determines how long fetchEncounters' cached responses get to exist before being discarded.
//   - description: Describes the purpose of the ExploreCommand.
//   - fetchEncounters: A function that fetches locations/maps from the Pokemon API.
//   - listMarker: A string marker to display before each map name for nicer formatting.
//   - listTitle: A string title to display before the list of maps.
//   - name: The name of the ExploreCommand.
//   - noEnteredArgsErrorMessage: An error message to display when no arguments are entered.
//   - noEncountersFoundErrorMessage: An error message to display when no maps are found.
//   - state: The state of the location area to be explored.
//
// Methods:
//   - Execute: Executes the ExploreCommand and fetches locations using the provided state and query.
//   - GetArgs: Returns the arguments of the ExploreCommand.
//   - GetDescription: Returns the description of the ExploreCommand.
//   - GetName: Returns the name of the ExploreCommand.
//   - PrintHelp: Prints the help message for the ExploreCommand.
//   - SetArgs: Sets the arguments of the ExploreCommand.
//   - hasValidArgs: Checks if the ExploreCommand has valid arguments needed to run the Execute() method properly.
type ExploreCommand struct {
	// The arguments provided to the command.
	args []string
	// The URL of the API endpoint to fetch the location area data.
	apiEndpoint string
	// cacheTTL is the time-to-live (TTL) determines how long FetchEncounters' cached responses get to exist before being discarded.
	cacheTTL time.Duration
	// description describes the purpose of the ExploreCommand.
	description string
	// fetchEncounters is a function that fetches locations/maps from the Pokemon API.
	fetchEncounters FetchEncounters
	// listMarker is a string marker to display before each map name for nicer formatting.
	listMarker string
	// listTitle is a string title to display before the list of maps.
	listTitle string
	// name is the name of the ExploreCommand.
	name string
	// noEncountersFoundErrorMessage is an error message to display when no maps are found.
	noEnteredArgsErrorMessage string
	// noEncountersFoundErrorMessage is an error message to display when no maps are found.
	noEncountersFoundErrorMessage string
	// state stores the fetched queried data.
	state *la.LocationArea
}

// Execute executes the MapCommand and fetches locations using the provided state and query.
//
// Returns:
//   - An error if the command execution fails.
//
// Example usage:
//
//	command := NewExploreCommand()
//	command.Execute()
func (e *ExploreCommand) Execute() error {
	if !e.hasValidArgs() {
		return fmt.Errorf(e.noEnteredArgsErrorMessage)
	}

	locationAreaEndpoint := e.apiEndpoint + "/" + e.args[0]
	if err := e.fetchEncounters(locationAreaEndpoint, e.state, e.cacheTTL); err != nil {
		return err
	}

	if len(e.state.Encounters) == 0 {
		return fmt.Errorf(e.noEncountersFoundErrorMessage)
	}

	fmt.Println(e.listTitle)
	for _, encounter := range e.state.Encounters {
		fmt.Printf("%s %s\n", e.listMarker, encounter.Pokemon.Name)
	}

	return nil
}

// GetArgs returns the arguments of the ExploreCommand.
//
// It will reset the args field value after getting its stored value.
//
// Returns:
//   - The arguments of the ExploreCommand.
//
// Example usage:
//
//	command := NewExploreCommand()
//	args := command.GetArgs()
func (e ExploreCommand) GetArgs() []string {
	defer e.SetArgs([]string{}) // Reset the args field value after getting thee.
	return e.args
}

// GetDescription returns the description of the MapCommand.
//
// Returns:
//   - The description of the MapCommand.
//
// Example usage:
//
//	command := NewExploreCommand()
//	description := command.GetDescription()
func (e ExploreCommand) GetDescription() string {
	return e.description
}

// GetName returns the name of the MapCommand.
//
// Returns:
//   - The name of the MapCommand.
//
// Example usage:
//
//	command := NewExploreCommand()
//	name := command.GetName()
func (e ExploreCommand) GetName() string {
	return e.name
}

// PrintHelp prints the help message for the MapCommand.
//
// Example usage:
//
//	command := NewExploreCommand()
//	command.PrintHelp()
func (e *ExploreCommand) PrintHelp() {
	c.PrintHelp(e)
}

// SetArgs sets the arguments of the ExploreCommand.
func (e *ExploreCommand) SetArgs(args []string) {
	e.args = args
}

// hasValidArgs checks if the ExploreCommand has valid arguments needed to run the Execute() method properly.
//
// Returns:
//   - A boolean indicating if the ExploreCommand has valid arguments.
//
// Example usage:
//
//	command := NewExploreCommand()
//	command.SetArgs([]string{"1"})
//	hasValidArgs := command.hasValidArgs()
func (e ExploreCommand) hasValidArgs() bool {
	return e.args != nil && len(e.args) > 0 && e.args[0] != ""
}
