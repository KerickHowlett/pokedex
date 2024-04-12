package pokemon_encounter

import p "explore/pokemon"

// WithPokemon sets the Pokemon of the PokemonEncounter.
//
// Parameters:
//   - pokemon: The Pokemon that can be encountered.
//
// Returns:
//   - PokemonEncounterOption: The option to set the Pokemon of the PokemonEncounter.
//
// Example:
//
//	pe := &PokemonEncounter{}
//	WithPokemon(p.Pokemon{Name: "pikachu"})(pe)
func WithPokemon(pokemon *p.Pokemon) PokemonEncounterOption {
	return func(pe *PokemonEncounter) {
		pe.Pokemon = pokemon
	}
}
