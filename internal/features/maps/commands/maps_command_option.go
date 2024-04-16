package maps_commands

import (
	"time"

	ms "maps/state"
)

type MapsCommandOption func(*MapsCommand)

// WithCacheTTL sets the time-to-live (TTL) for the cache.
//
// Parameters:
//   - cacheTTL: A time.Duration containing the time-to-live (TTL) for the cache.
//
// Returns:
//   - A MapsCommandOption function that sets the time-to-live (TTL) for the cache.
//
// Example usage:
//
//	command := NewMapsCommand(WithCacheTTL(ttl.OneDayTTL))
func WithCacheTTL(cacheTTL time.Duration) MapsCommandOption {
	return func(c *MapsCommand) {
		c.cacheTTL = cacheTTL
	}
}

// WithCommandDescription sets the description of the command.
//
// Parameters:
//   - commandDescription: A string containing the description of the command.
//
// Returns:
//   - A MapsCommandOption function that sets the description of the command.
//
// Example usage:
//
//	command := NewMapsCommand(WithCommandDescription("Command Description"))
func WithCommandDescription(commandDescription string) MapsCommandOption {
	return func(c *MapsCommand) {
		c.description = commandDescription
	}
}

// WithCommandName sets the name of the command.
//
// Parameters:
//   - commandName: A string containing the name of the command.
//
// Returns:
//   - A MapsCommandOption function that sets the name of the command.
//
// Example usage:
//
//	command := NewMapsCommand(WithCommandName("map"))
func WithCommandName(commandName string) MapsCommandOption {
	return func(c *MapsCommand) {
		c.name = commandName
	}
}

// WithFetchLocations sets the function that fetches the locations.
//
// Parameters:
//   - fetchLocations: A function that fetches the locations.
//
// Returns:
//   - A MapsCommandOption function that sets the function that fetches the locations.
//
// Example usage:
//
//	command := NewMapsCommand(WithFetchLocations(qf.QueryFetch[ms.MapsState]))
func WithFetchLocations(fetchLocations FetchLocations) MapsCommandOption {
	return func(c *MapsCommand) {
		c.fetchLocations = fetchLocations
	}
}

// WithListMarker sets the marker to display before each map name.
//
// Parameters:
//   - listMarker: A string containing the marker to display before each map name.
//
// Returns:
//   - A MapsCommandOption function that sets the marker to display before each map name.
//
// Example usage:
//
//	command := NewMapsCommand(WithListMarker("  -"))
func WithListMarker(listMarker string) MapsCommandOption {
	return func(c *MapsCommand) {
		c.listMarker = listMarker
	}
}

// WithListTitle sets the title to display before the list of maps.
//
// Parameters:
//   - listTitle: A string containing the title to display before the list of maps.
//
// Returns:
//   - A MapsCommandOption function that sets the title to display before the list of maps.
//
// Example usage:
//
//	command := NewMapsCommand(WithListTitle("Pokemon Maps:"))
func WithListTitle(listTitle string) MapsCommandOption {
	return func(c *MapsCommand) {
		c.listTitle = listTitle
	}
}

// WithNoMapsFoundErrorMessage sets the error message to display when no maps are found.
//
// Parameters:
//   - noMapsFoundErrorMessage: A string containing the error message to display when no maps are found.
//
// Returns:
//   - A MapsCommandOption function that sets the error message to display when no maps are found.
//
// Example usage:
//
//	command := NewMapsCommand(WithNoMapsFoundErrorMessage("no maps were found"))
func WithNoMapsFoundErrorMessage(noMapsFoundErrorMessage string) MapsCommandOption {
	return func(c *MapsCommand) {
		c.noMapsFoundErrorMessage = noMapsFoundErrorMessage
	}
}

// WithNoMoreMapsMessage sets the message to display when there are no more maps.
//
// Parameters:
//   - noMoreMapsMessage: A string containing the message to display when there are no more maps.
//
// Returns:
//   - A MapsCommandOption function that sets the message to display when there are no more maps.
//
// Example usage:
//
//	command := NewMapsCommand(WithNoMoreMapsMessage("There are no more maps."))
func WithNoMoreMapsMessage(noMoreMapsMessage string) MapsCommandOption {
	return func(c *MapsCommand) {
		c.noMoreMapsMessage = noMoreMapsMessage
	}
}

// WithPaginationDirection sets the direction of the maps pagination.
//
// Parameters:
//   - paginationDirection: A string containing the direction of the maps pagination.
//
// Returns:
//   - A MapsCommandOption function that sets the direction of the maps pagination.
//
// Example usage:
//
//	command := NewMapsCommand(WithPaginationDirection("next"))
func WithPaginationDirection(paginationDirection string) MapsCommandOption {
	return func(c *MapsCommand) {
		c.paginationDirection = paginationDirection
	}
}

// WithState sets the state of the command.
//
// Parameters:
//   - state: A pointer to a MapsState struct instance.
//
// Returns:
//   - A MapsCommandOption function that sets the state of the command.
//
// Example usage:
//
//	command := NewMapsCommand(WithState(ms.NewMapsState()))
func WithState(state *ms.MapsState) MapsCommandOption {
	return func(c *MapsCommand) {
		c.state = state
	}
}
