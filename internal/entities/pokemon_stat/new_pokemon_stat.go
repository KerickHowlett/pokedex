package pokemon_stat

import el "entity_link"

// NewPokemonStat creates a new PokemonStat with the provided options.
//
// Parameters:
//   - options: The options to set on the PokemonStat.
//
// Returns:
//   - *PokemonStat: The new PokemonStat.
//
// Example usage:
//
//	pokemonStat := NewPokemonStat(WithBaseStat(100), WithStat(el.NewEntityLink()))
func NewPokemonStat(options ...PokemonStatOption) *PokemonStat {
	pokemonStat := &PokemonStat{Stat: el.NewEntityLink()}
	for _, option := range options {
		option(pokemonStat)
	}
	return pokemonStat
}
