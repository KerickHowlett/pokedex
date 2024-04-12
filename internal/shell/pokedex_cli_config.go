package pokedexclishell

// PokedexCLIConfig represents the configuration for the Pokedex CLI.
//
// Fields:
//   - LocalAreaAPIEndpoint: The API endpoint for the Local Area service.
//   - Prompt: The text displayed as the command prompt.
//   - StartingMapsAPIEndpoint: The initial API endpoint for the Maps service.
type PokedexCLIConfig struct {
	LocalAreaAPIEndpoint    string
	Prompt                  string
	StartingMapsAPIEndpoint string
}
