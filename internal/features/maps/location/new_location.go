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
	pokemonMap := &Location{
		Name: "",
	}

	for _, option := range options {
		option(pokemonMap)
	}

	if pokemonMap.Name == "" {
		panic("[ERROR] Location name is required")
	}

	return pokemonMap
}
