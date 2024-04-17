package pokemon

// NewPokemon creates a new Pokemon with the provided options.
// It returns a pointer to the created Pokemon.
// If the name option is not provided, it panics with an error message.
//
// Parameters:
//   - options: A variadic parameter of PokemonOption functions.
//
// Returns:
//   - A pointer to the created Pokemon.
//
// Example usage:
//
//	pokemon := NewPokemon()
//	fmt.Println(pokemon.Name) // Output: Pikachu
func NewPokemon(options ...PokemonOption) *Pokemon {
	pokemonMap := &Pokemon{}
	for _, option := range options {
		option(pokemonMap)
	}
	return pokemonMap
}
