package pokemon_stat

import el "entity_link"

// PokemonStat represents a single stat name & value for a Pokemon.
//
// Fields:
//   - BaseStat: The value associated with the owning Pokemon's stat.
//   - Stat: The name of the stat and the link to its associated entity information.
type PokemonStat struct {
	// The value associated with the owning Pokemon's stat.
	BaseStat int `json:"base_stat"`
	// The name of the stat and the link to its associated entity information.
	Stat *el.EntityLink `json:"stat"`
}
