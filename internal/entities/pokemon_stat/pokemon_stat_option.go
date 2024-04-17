package pokemon_stat

import el "entity_link"

// PokemonStatOption represents a single stat name & value for a Pokemon.
type PokemonStatOption func(*PokemonStat)

// WithBaseStat sets the value associated with the owning Pokemon's stat.
//
// Parameters:
//   - baseStat: The value associated with the owning Pokemon's stat.
//
// Example usage:
//
//	pokemonStat := NewPokemonStat(WithBaseStat(100))
func WithBaseStat(baseStat int) PokemonStatOption {
	return func(ps *PokemonStat) {
		ps.BaseStat = baseStat
	}
}

// WithStat sets the name of the stat and the link to its associated entity information.
//
// Parameters:
//   - stat: The name of the stat and the link to its associated entity information.
//
// Example usage:
//
//	pokemonStat := NewPokemonStat(WithStat(el.NewEntityLink()))
func WithStat(stat *el.EntityLink) PokemonStatOption {
	return func(ps *PokemonStat) {
		ps.Stat = stat
	}
}
