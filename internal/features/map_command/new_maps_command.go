package map_command

import (
	"fmt"
	ms "map_command/state"
	qf "query/fetch"
	"query/fetch/ttl"
)

// NewMapCommand creates a new instance of the Map struct.
//
// Parameters:
//   - options: A variadic parameter of MapCommandOption functions.
//
// Returns:
//   - A new instance of the Map struct.
//
// Example usage:
//
//	command := NewMapCommand()
func NewMapCommand(options ...MapCommandOption) (command *MapCommand) {
	command = &MapCommand{
		cacheTTL:                ttl.OneDayTTL,
		fetchLocations:          qf.QueryFetch[ms.MapsState],
		listMarker:              "  -",
		listTitle:               "Pokemon Maps:",
		noMapsFoundErrorMessage: "no maps were found",
		state:                   ms.NewMapsState(),
	}

	for _, option := range options {
		option(command)
	}

	if command.paginationDirection == "" {
		panic("a pagination direction option of 'next' or 'previous' is required for the maps command")
	}

	if command.name == "" {
		command.name = setDefaultCommandName(command.paginationDirection)
	}

	if command.description == "" {
		command.description = fmt.Sprintf("Show the %s 20 Pokemon world map locations.", command.paginationDirection)
	}

	if command.noMoreMapsMessage == "" {
		command.noMoreMapsMessage = setDefaultNoMoreMapsMessage(command.paginationDirection)
	}

	return command
}

// setDefaultCommandName sets the default command name based on the pagination direction.
//
// This is a private function that is used within the NewMapCommand function.
//
// Parameters:
//   - paginationDirection: The direction of the pagination. Either "next" or "previous".
//
// Returns:
//   - The default command name.
//   - Panics if the pagination direction argument is invalid.
//
// Example usage:
//
// command.commandName = setDefaultCommandName(Next)
func setDefaultCommandName(paginationDirection string) string {
	switch paginationDirection {
	case "next":
		return "map"
	case "previous":
		return "mapb"
	default:
		panic("Invalid pagination direction")
	}
}

// setDefaultNoMoreMapsMessage sets the default message when there are no more maps to display.
//
// This is a private function that is used within the NewMapCommand function.
//
// Parameters:
//   - paginationDirection: The direction of the pagination. Either "next" or "previous".
//
// Returns:
//   - The default message when there are no more maps to display.
//   - Panics if the pagination direction argument is invalid.
//
// Example usage:
//
// command.noMoreMapsMessage = setDefaultNoMoreMapsMessage(Next)
func setDefaultNoMoreMapsMessage(paginationDirection string) string {
	switch paginationDirection {
	case Next:
		return "There are no more maps."
	case Previous:
		return "You're already at the beginning of the maps list."
	default:
		panic("Invalid pagination direction")
	}
}
