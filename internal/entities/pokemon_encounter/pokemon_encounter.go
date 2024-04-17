package pokemon_encounter

import el "entity_link"

// PokemonEncounter houses the essential information of a given Pokemon encounter.
//
// Fields:
//   - Pokemon: The Pokemon that can be encountered.
type PokemonEncounter struct {
	// Pokemon represents the Pokemon that can be encountered.
	Pokemon *el.EntityLink `json:"pokemon"`
}
