package pokedexclishell

// PokedexCLIConfig represents the configuration for the Pokedex CLI.
//
// Fields:
//   - LocalAreaAPIEndpoint: The API endpoint for the Local Area service.
//   - PokemonAPIEndpoint: The API endpoint for the Pokemon service.
//   - Prompt: The text displayed as the command prompt.
//   - StartingMapsAPIEndpoint: The initial API endpoint for the Maps service.
type PokedexCLIConfig struct {
	// LocalAreaAPIEndpoint is the API endpoint for the Local Area service.
	LocalAreaAPIEndpoint string
	// PokemonAPIEndpoint is the API endpoint for the Pokemon service.
	PokemonAPIEndpoint string
	// Prompt is the text displayed as the command prompt.
	Prompt string
	// StartingMapsAPIEndpoint is the initial API endpoint for the Maps service.
	StartingMapsAPIEndpoint string
}
