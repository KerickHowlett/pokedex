package explore

import la "explore/location_area"

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
		apiEndpoint: "",
		state:       la.NewLocationArea(),
	}

	for _, option := range options {
		option(command)
	}

	if command.apiEndpoint == "" {
		panic("API endpoint is required")
	}

	return command
}
