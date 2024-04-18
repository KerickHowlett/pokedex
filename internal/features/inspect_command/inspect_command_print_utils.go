package inspect_command

import (
	"fmt"

	p "pokemon"
)

// printBasicInfoOf prints the basic information of a Pokemon.
//
// Parameters:
//   - pokemon: The Pokemon to print the basic information of.
//
// Example usage:
//
//	pokemon := &p.Pokemon{Name: "Pikachu", Height: 4, Weight: 60}
//	printBasicInfoOf(pokemon)
func printBasicInfoOf(pokemon *p.Pokemon) {
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
}

// printTypesOf prints the types of a Pokemon.
//
// Parameters:
//   - pokemon: The Pokemon to print the types of.
//
// Example usage:
//
//	pokemon := &p.Pokemon{Types: []p.PokemonType{{Type: p.Type{Name: "Electric"}}}}
//	printTypesOf(pokemon)
func printTypesOf(pokemon *p.Pokemon) {
	fmt.Println("Types:")
	for _, pokemonType := range pokemon.Types {
		fmt.Printf(" - %s\n", pokemonType.Type.Name)
	}
}

// printStatsOf prints the stats of a Pokemon.
//
// Parameters:
//   - pokemon: The Pokemon to print the stats of.
//
// Example usage:
//
//	pokemon := &p.Pokemon{Stats: []p.PokemonStat{{Stat: p.Stat{Name: "hp"}, BaseStat: 35}}}
//	printStatsOf(pokemon)
func printStatsOf(pokemon *p.Pokemon) {
	fmt.Println("Stats:")
	for _, pokemonStat := range pokemon.Stats {
		fmt.Printf(" - %s: %d\n", pokemonStat.Stat.Name, pokemonStat.BaseStat)
	}
}
