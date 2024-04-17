package pokemon_encounter

import el "entity_link"

type PokemonEncounterOption func(*PokemonEncounter)

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
//	WithPokemon(el.EntityLink{Name: "pikachu"})(pe)
func WithPokemon(pokemon *el.EntityLink) PokemonEncounterOption {
	return func(pe *PokemonEncounter) {
		pe.Pokemon = pokemon
	}
}
