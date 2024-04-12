package location

// NewLocation creates a new Location with the provided options.
// It returns a pointer to the created Location.
// If the name option is not provided, it panics with an error message.
//
// Parameters:
//   - options: A variadic parameter of LocationOption functions.
//
// Returns:
//   - A pointer to the created Location.
//
// Example usage:
//
//	location := NewLocation(
//		WithName("Pallet Town"),
//	)
//	fmt.Println(location.Name) // Output: Pallet Town
func NewLocation(options ...LocationOption) *Location {
	location := &Location{}

	for _, option := range options {
		option(location)
	}

	if location.Name == "" {
		panic("name of location is required")
	}

	return location
}
