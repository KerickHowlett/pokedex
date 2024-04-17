package pokemon

import (
	ps "pokemon_stat"
	pt "pokemon_type"
)

// Pokemon houses the essential information of a given Pokemon.
//
// Fields:
//   - BaseExperience: The base experience of the Pokemon, which effects the catch rate.
//   - Height: The height of the Pokemon.
//   - Name: The name of the Pokemon.
//   - Stats: The stats of the Pokemon (e.g., HP, attack, defense).
//   - Types: The types of the Pokemon (e.g., fire, grass, psychic).
//   - Weight: The weight of the Pokemon.
type Pokemon struct {
	// BaseExperience represents the base experience of the Pokemon, which effects the catch rate.
	BaseExperience int `json:"base_experience"`
	// Height represents the height of the Pokemon.
	Height int `json:"height"`
	// Name represents the name of the Pokemon.
	Name string `json:"name"`
	// Stats represents the stats of the Pokemon (e.g., HP, attack, defense).
	Stats []*ps.PokemonStat `json:"stats"`
	// Types represents the types of the Pokemon (e.g., fire, grass, psychic).
	Types []*pt.PokemonType `json:"types"`
	// Weight represents the weight of the Pokemon.
	Weight int `json:"weight"`
}
