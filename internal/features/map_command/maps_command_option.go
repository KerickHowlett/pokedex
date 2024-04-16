package map_command

import (
	"time"

	ms "map_command/state"
)

type MapCommandOption func(*MapCommand)

// WithCacheTTL sets the time-to-live (TTL) for the cache.
//
// Parameters:
//   - cacheTTL: A time.Duration containing the time-to-live (TTL) for the cache.
//
// Returns:
//   - A MapCommandOption function that sets the time-to-live (TTL) for the cache.
//
// Example usage:
//
//	command := NewMapCommand(WithCacheTTL(ttl.OneDayTTL))
func WithCacheTTL(cacheTTL time.Duration) MapCommandOption {
	return func(c *MapCommand) {
		c.cacheTTL = cacheTTL
	}
}

// WithCommandDescription sets the description of the command.
//
// Parameters:
//   - commandDescription: A string containing the description of the command.
//
// Returns:
//   - A MapCommandOption function that sets the description of the command.
//
// Example usage:
//
//	command := NewMapCommand(WithCommandDescription("Command Description"))
func WithCommandDescription(commandDescription string) MapCommandOption {
	return func(c *MapCommand) {
		c.description = commandDescription
	}
}

// WithCommandName sets the name of the command.
//
// Parameters:
//   - commandName: A string containing the name of the command.
//
// Returns:
//   - A MapCommandOption function that sets the name of the command.
//
// Example usage:
//
//	command := NewMapCommand(WithCommandName("map"))
func WithCommandName(commandName string) MapCommandOption {
	return func(c *MapCommand) {
		c.name = commandName
	}
}

// WithFetchLocations sets the function that fetches the locations.
//
// Parameters:
//   - fetchLocations: A function that fetches the locations.
//
// Returns:
//   - A MapCommandOption function that sets the function that fetches the locations.
//
// Example usage:
//
//	command := NewMapCommand(WithFetchLocations(qf.QueryFetch[ms.MapsState]))
func WithFetchLocations(fetchLocations FetchLocations) MapCommandOption {
	return func(c *MapCommand) {
		c.fetchLocations = fetchLocations
	}
}

// WithListMarker sets the marker to display before each map name.
//
// Parameters:
//   - listMarker: A string containing the marker to display before each map name.
//
// Returns:
//   - A MapCommandOption function that sets the marker to display before each map name.
//
// Example usage:
//
//	command := NewMapCommand(WithListMarker("  -"))
func WithListMarker(listMarker string) MapCommandOption {
	return func(c *MapCommand) {
		c.listMarker = listMarker
	}
}

// WithListTitle sets the title to display before the list of maps.
//
// Parameters:
//   - listTitle: A string containing the title to display before the list of maps.
//
// Returns:
//   - A MapCommandOption function that sets the title to display before the list of maps.
//
// Example usage:
//
//	command := NewMapCommand(WithListTitle("Pokemon Maps:"))
func WithListTitle(listTitle string) MapCommandOption {
	return func(c *MapCommand) {
		c.listTitle = listTitle
	}
}

// WithNoMapsFoundErrorMessage sets the error message to display when no maps are found.
//
// Parameters:
//   - noMapsFoundErrorMessage: A string containing the error message to display when no maps are found.
//
// Returns:
//   - A MapCommandOption function that sets the error message to display when no maps are found.
//
// Example usage:
//
//	command := NewMapCommand(WithNoMapsFoundErrorMessage("no maps were found"))
func WithNoMapsFoundErrorMessage(noMapsFoundErrorMessage string) MapCommandOption {
	return func(c *MapCommand) {
		c.noMapsFoundErrorMessage = noMapsFoundErrorMessage
	}
}

// WithNoMoreMapsMessage sets the message to display when there are no more maps.
//
// Parameters:
//   - noMoreMapsMessage: A string containing the message to display when there are no more maps.
//
// Returns:
//   - A MapCommandOption function that sets the message to display when there are no more maps.
//
// Example usage:
//
//	command := NewMapCommand(WithNoMoreMapsMessage("There are no more maps."))
func WithNoMoreMapsMessage(noMoreMapsMessage string) MapCommandOption {
	return func(c *MapCommand) {
		c.noMoreMapsMessage = noMoreMapsMessage
	}
}

// WithPaginationDirection sets the direction of the maps pagination.
//
// Parameters:
//   - paginationDirection: A string containing the direction of the maps pagination.
//
// Returns:
//   - A MapCommandOption function that sets the direction of the maps pagination.
//
// Example usage:
//
//	command := NewMapCommand(WithPaginationDirection("next"))
func WithPaginationDirection(paginationDirection string) MapCommandOption {
	return func(c *MapCommand) {
		c.paginationDirection = paginationDirection
	}
}

// WithState sets the state of the command.
//
// Parameters:
//   - state: A pointer to a MapsState struct instance.
//
// Returns:
//   - A MapCommandOption function that sets the state of the command.
//
// Example usage:
//
//	command := NewMapCommand(WithState(ms.NewMapsState()))
func WithState(state *ms.MapsState) MapCommandOption {
	return func(c *MapCommand) {
		c.state = state
	}
}
