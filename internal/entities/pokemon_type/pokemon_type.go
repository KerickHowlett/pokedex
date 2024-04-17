package pokemon_type

import el "entity_link"

// PokemonTypeOption represents the functional option pattern for PokemonType.
//
// Fields:
// - Slot: the slot of the PokemonType (i.e., Primary or Secondary).
// - Type: link to the Pokemon Type's entity information.
type PokemonType struct {
	// Slot represents the slot of the PokemonType (i.e., Primary or Secondary).
	Slot int `json:"slot"`
	// Type represents link to the Pokemon Type's entity information.
	Type *el.EntityLink `json:"type"`
}
