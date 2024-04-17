package pokemon

import (
	ps "pokemon_stat"
	pt "pokemon_type"
)

// PokemonOption is a function that modifies a Pokemon struct.
type PokemonOption func(*Pokemon)

// WithBaseExperience returns a function that sets the
// base experience of a Pokemon.
//
// Parameters:
//   - baseExperience: The base experience of the Pokemon.
//
// Returns:
//   - A PokemonOption function.
//
// Example usage:
//
//	pokemon := NewPokemon(
//		WithBaseExperience(100),
//	)
func WithBaseExperience(baseExperience int) PokemonOption {
	return func(p *Pokemon) {
		p.BaseExperience = baseExperience
	}
}

// WithHeight returns a function that sets the
// height of a Pokemon.
//
// Parameters:
//   - height: The height of the Pokemon.
//
// Returns:
//   - A PokemonOption function.
//
// Example usage:
//
//	pokemon := NewPokemon(
//		WithHeight(100),
//	)
func WithHeight(height int) PokemonOption {
	return func(p *Pokemon) {
		p.Height = height
	}
}

// WithName returns a function that sets the
// name of a Pokemon.
//
// The returned function can be used as a
// modifier function to set the name of a
// Pokemon struct.
//
// Parameters:
//   - name: The name of the Pokemon.
//
// Returns:
//   - A PokemonOption function.
//
// Example usage:
//
//	pokemon := NewPokemon(
//		WithName("Pikachu"),
//	)
//	fmt.Println(pokemon.Name) // Output: Pikachu
func WithName(name string) PokemonOption {
	return func(p *Pokemon) {
		p.Name = name
	}
}

// WithWeight returns a function that sets the
// weight of a Pokemon.
//
// Parameters:
//   - weight: The weight of the Pokemon.
//
// Returns:
//   - A PokemonOption function.
//
// Example usage:
//
//	pokemon := NewPokemon(
//		WithWeight(100),
//	)
func WithWeight(weight int) PokemonOption {
	return func(p *Pokemon) {
		p.Weight = weight
	}
}

// WithStat adds a stat onto a given pokemon.
//
// Parameters:
//   - stat: The stat to add to the Pokemon.
//
// Returns:
//   - A PokemonOption function.
//
// Example usage:
//
//	pokemon := NewPokemon(
//		WithStat(*pokemonStat),
//	)
func WithStat(stat *ps.PokemonStat) PokemonOption {
	return func(p *Pokemon) {
		p.Stats = append(p.Stats, stat)
	}
}

// WithType adds a type onto a given pokemon.
//
// Parameters:
//   - type: The type to add to the Pokemon.
//
// Returns:
//   - A PokemonOption function.
//
// Example usage:
//
//	pokemon := NewPokemon(
//		WithType(*pokemonType),
//	)
func WithType(pokemonType *pt.PokemonType) PokemonOption {
	return func(p *Pokemon) {
		p.Types = append(p.Types, pokemonType)
	}
}
