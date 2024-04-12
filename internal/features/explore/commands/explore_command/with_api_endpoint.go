package explore

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
func WithAPIEndpoint(apiEndpoint string) func(*ExploreCommand) {
	return func(e *ExploreCommand) {
		e.apiEndpoint = apiEndpoint
	}
}
