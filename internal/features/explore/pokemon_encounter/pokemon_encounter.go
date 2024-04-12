package pokemon_encounter

import p "explore/pokemon"

// PokemonEncounter houses the essential information of a given Pokemon encounter.
//
// Fields:
//   - Pokemon: The Pokemon that can be encountered.
type PokemonEncounter struct {
	// Pokemon represents the Pokemon that can be encountered.
	Pokemon *p.Pokemon `json:"pokemon"`
}
