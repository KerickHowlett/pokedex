package explore_command

import qec "query_fetch/query_cache/cache_eviction_config"

type ExploreCommandOption func(*ExploreCommand)

// WithAPIEndpoint sets the API endpoint for the ExploreCommand instance.
//
// Parameters:
//   - apiEndpoint: A string representing the API endpoint.
//
// Returns:
//   - A new ExploreCommandOption instance.
//
// Example usage:
//
//	command := NewExploreCommand(WithAPIEndpoint("http://localhost:8080"))
func WithAPIEndpoint(apiEndpoint string) ExploreCommandOption {
	return func(e *ExploreCommand) {
		e.apiEndpoint = apiEndpoint
	}
}

// WithDescription sets the description for the ExploreCommand instance.
//
// Parameters:
//   - description: A string representing the description.
//
// Returns:
//   - A new ExploreCommandOption instance.
//
// Example usage:
//
//	command := NewExploreCommand(WithDescription("Command Description"))
func WithDescription(description string) ExploreCommandOption {
	return func(e *ExploreCommand) {
		e.description = description
	}
}

// WithEvictionConfig sets the eviction configuration for the ExploreCommand instance.
//
// Parameters:
//   - ec: A pointer to a QueryEvictionConfig instance.
//
// Returns:
//   - A new ExploreCommandOption instance.
//
// Example usage:
//
//	command := NewExploreCommand(WithEvictionConfig(qec.NewQueryEvictionConfig()))
func WithEvictionConfig(ec *qec.QueryEvictionConfig) ExploreCommandOption {
	return func(e *ExploreCommand) {
		e.ec = ec
	}
}

// WithFetchEncounters sets the fetchEncounters function for the ExploreCommand instance.
//
// Parameters:
//   - fetchEncounters: A function that fetches encounters.
//
// Returns:
//   - A new ExploreCommandOption instance.
//
// Example usage:
//
//	command := NewExploreCommand(WithFetchEncounters(fetchEncounters))
func WithFetchEncounters(fetchEncounters FetchEncounters) ExploreCommandOption {
	return func(e *ExploreCommand) {
		e.fetchEncounters = fetchEncounters
	}
}

// WithListMarker sets the list marker for the ExploreCommand instance.
//
// Parameters:
//   - listMarker: A string representing the list marker.
//
// Returns:
//   - A new ExploreCommandOption instance.
//
// Example usage:
//
//	command := NewExploreCommand(WithListMarker(" -"))
func WithListMarker(listMarker string) ExploreCommandOption {
	return func(e *ExploreCommand) {
		e.listMarker = listMarker
	}
}

// WithListTitle sets the list title for the ExploreCommand instance.
//
// Parameters:
//   - listTitle: A string representing the list title.
//
// Returns:
//   - A new ExploreCommandOption instance.
//
// Example usage:
//
//	command := NewExploreCommand(WithListTitle("List Title"))
func WithListTitle(listTitle string) ExploreCommandOption {
	return func(e *ExploreCommand) {
		e.listTitle = listTitle
	}
}

// WithName sets the name for the ExploreCommand instance.
//
// Parameters:
//   - name: A string representing the name.
//
// Returns:
//   - A new ExploreCommandOption instance.
//
// Example usage:
//
//	command := NewExploreCommand(WithName("Command Name"))
func WithName(name string) ExploreCommandOption {
	return func(e *ExploreCommand) {
		e.name = name
	}
}

// WithNoEncountersFoundErrorMessage sets the no encounters found error message for the ExploreCommand instance.
//
// Parameters:
//   - noEncountersFoundErrorMessage: A string representing the no encounters found error message.
//
// Returns:
//   - A new ExploreCommandOption instance.
//
// Example usage:
//
//	command := NewExploreCommand(WithNoEncountersFoundErrorMessage("No encounters found error message"))
func WithNoEncountersFoundErrorMessage(noEncountersFoundErrorMessage string) ExploreCommandOption {
	return func(e *ExploreCommand) {
		e.noEncountersFoundErrorMessage = noEncountersFoundErrorMessage
	}
}

// WithNoEnteredArgsErrorMessage sets the no entered arguments error message for the ExploreCommand instance.
//
// Parameters:
//   - noEnteredArgsErrorMessage: A string representing the no entered arguments error message.
//
// Returns:
//   - A new ExploreCommandOption instance.
//
// Example usage:
//
//	command := NewExploreCommand(WithNoEnteredArgsErrorMessage("No entered arguments error message"))
func WithNoEnteredArgsErrorMessage(noEnteredArgsErrorMessage string) ExploreCommandOption {
	return func(e *ExploreCommand) {
		e.noEnteredArgsErrorMessage = noEnteredArgsErrorMessage
	}
}
