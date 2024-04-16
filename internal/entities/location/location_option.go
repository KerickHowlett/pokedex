package location

// LocationOption is a function that modifies a Location struct.
type LocationOption func(*Location)

// WithName returns a function that sets the
// name of a Location.
//
// The returned function can be used as a
// modifier function to set the name of a
// Location struct.
//
// Parameters:
//   - name: The name of the Location.
//
// Returns:
//   - A LocationOption function.
//
// Example usage:
//
//	location := NewLocation(
//		WithName("Pallet Town"),
//	)
//	fmt.Println(location.Name) // Output: Pallet Town
func WithName(name string) func(*Location) {
	return func(p *Location) {
		p.Name = name
	}
}
