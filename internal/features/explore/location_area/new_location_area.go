package location_area

// LocationAreaOption is a function that modifies a LocationArea
// struct.
//
// Parameters:
//   - options: A variadic list of LocationAreaOption functions used to modify the LocationArea struct.
//
// Returns:
//   - A new LocationArea struct instance.
//
// Example usage:
//
//	locationArea := NewLocationArea()
func NewLocationArea(options ...LocationAreaOption) LocationArea {
	locationArea := LocationArea{}
	for _, option := range options {
		option(&locationArea)
	}
	return locationArea
}
