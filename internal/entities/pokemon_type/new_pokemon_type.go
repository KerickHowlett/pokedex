package pokemon_type

import el "entity_link"

func NewPokemonType(options ...PokemonTypeOption) *PokemonType {
	pokemonType := &PokemonType{Type: el.NewEntityLink()}
	for _, option := range options {
		option(pokemonType)
	}
	return pokemonType
}
