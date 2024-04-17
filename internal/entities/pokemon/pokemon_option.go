package pokemon

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
func WithBaseExperience(baseExperience int) func(*Pokemon) {
	return func(p *Pokemon) {
		p.BaseExperience = baseExperience
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
func WithName(name string) func(*Pokemon) {
	return func(p *Pokemon) {
		p.Name = name
	}
}
