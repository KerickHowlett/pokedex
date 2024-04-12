package location_area

import pe "entities/pokemon_encounter"

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
	return func(p *LocationArea) {
		p.Encounters = append(p.Encounters, encounter)
	}
}
