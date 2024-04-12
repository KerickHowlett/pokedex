package pokemon_encounter

// PokemonEncounter houses the essential information of a given Pokemon encounter.
//
// Parameters:
//   - Pokemon: The Pokemon that can be encountered.
//
// Returns:
//   - A new PokemonEncounter struct instance set with the provided options.
//
// Example:
//
//	pe := NewPokemonEncounter(WithPokemon(p.Pokemon{Name: "pikachu"}))
func NewPokemonEncounter(options ...PokemonEncounterOption) *PokemonEncounter {
	pe := &PokemonEncounter{}
	for _, option := range options {
		option(pe)
	}
	return pe
}
