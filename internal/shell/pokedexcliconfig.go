package pokedexclishell

// PokedexCLIConfig represents the configuration for the Pokedex CLI.
//
// Fields:
//   - Prompt: The text displayed as the command prompt.
//   - StartingMapsAPIEndpoint: The initial API endpoint for the Maps service.
type PokedexCLIConfig struct {
	Prompt                  string
	StartingMapsAPIEndpoint string
}
