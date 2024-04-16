package location_area

import pe "pokemon_encounter"

// LocationAreaOption is a function that modifies a LocationArea struct.
type LocationAreaOption func(*LocationArea)

// WithPokemonEncounter returns a function that sets the
// Pokemon encounters of a LocationArea.
//
// The returned function can be used as a
// modifier function to set the Pokemon encounters of a
// LocationArea struct.
//
// Parameters:
//   - encounter: The Pokemon that can be encounter within the LocationArea.
//
// Returns:
//   - A LocationAreaOption function.
//
// Example usage:
//
//	locationArea := NewLocationArea(
//		WithPokemonEncounter(pokemon),
//	)
func WithPokemonEncounter(encounter pe.PokemonEncounter) LocationAreaOption {
	return func(la *LocationArea) {
		la.Encounters = append(la.Encounters, encounter)
	}
}
