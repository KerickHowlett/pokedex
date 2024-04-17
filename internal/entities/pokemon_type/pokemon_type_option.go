package pokemon_type

import el "entity_link"

// PokemonTypeOption represents the functional option pattern for PokemonType.
type PokemonTypeOption func(*PokemonType)

// WithSlot sets the slot of the PokemonType.
//
// Parameters:
// - slot: the slot of the PokemonType.
//
// Returns:
// - PokemonTypeOption: the functional option to set the slot of the PokemonType.
//
// Example usage:
//
// pokemonType := NewPokemonType(WithSlot(1))
func WithSlot(slot int) PokemonTypeOption {
	return func(pt *PokemonType) {
		pt.Slot = slot
	}
}

// WithType sets the type of the PokemonType.
//
// Parameters:
// - pokemonType: the type of the PokemonType.
//
// Returns:
// - PokemonTypeOption: the functional option to set the type of the PokemonType.
//
// Example usage:
//
// pokemonType := NewPokemonType(WithType(el.NewEntityLink(pokemonType)))
func WithType(pokemonType *el.EntityLink) PokemonTypeOption {
	return func(pt *PokemonType) {
		pt.Type = pokemonType
	}
}
