package explore_command

import (
	la "entities/location_area"
	qf "query_fetch"
	"query_fetch/query_cache/ttl"
)

// NewExploreCommand creates a new instance of the Map struct.
//
// Parameters:
//   - state: A pointer to a LocationArea struct instance.
//
// Returns:
//   - A new instance of the Map struct.
//
// Example usage:
//
//	command := NewExploreCommand()
func NewExploreCommand(options ...ExploreCommandOption) *ExploreCommand {
	command := &ExploreCommand{
		cacheTTL:                      ttl.OneDayTTL,
		description:                   "Find out what Pokemon can be encountered in any location area. Usage: explore <location-area-name>",
		fetchEncounters:               qf.QueryFetch[la.LocationArea],
		listMarker:                    " -",
		listTitle:                     "Found Pokemon:",
		name:                          "explore",
		noEnteredArgsErrorMessage:     "a location area name is required",
		noEncountersFoundErrorMessage: "no pokemon encounters can be found",
		state:                         la.NewLocationArea(),
	}

	for _, option := range options {
		option(command)
	}

	if command.apiEndpoint == "" {
		panic("an api endpoint is required")
	}

	return command
}
