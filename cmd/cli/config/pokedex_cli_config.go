package config

import s "shell"

// @TODO: Implement means to pull certain property values from dotenv files
//        based on the environment the application is running in.
//        e.g., dev, qa, prod, etc.

// PokedexCLIConfig is a struct that holds the configuration for the Pokedex CLI.
//
// Configurations:
//   - ListSize: The number of items to display in a list.
//   - Offset: The starting index of the list.
//   - PokedexAPIDomain: The domain of the Pokedex API.
//   - PokedexPrompt: The text displayed as the Pokedex prompt.
//   - PokemonMapsURI: The URI for the Pokemon Maps service.
//   - Prompt: The text displayed as the command prompt.
//   - StartingMapsAPIEndpoint: The initial API endpoint for the Maps service.
//   - limitParam: The query parameter for the number of items to display in a list.
//   - offsetParam: The query parameter for the starting index of the list.
const (
	ListSize         = "20"
	Offset           = "0"
	PokedexAPIDomain = "https://pokeapi.co/api/v2"
	PokedexPrompt    = "pokedex > "
	PokemonMapsURI   = "/location"

	// Private Constants to be used for URI query parameters.
	offsetParam = "offset"
	limitParam  = "limit"
)

// StartingMapsAPIEndpoint is the initial API endpoint for the Maps service.
var (
	StartingMapsAPIEndpoint = PokedexAPIDomain + PokemonMapsURI + "?" + offsetParam + "=" + Offset + "&" + limitParam + "=" + ListSize
)

// PokedexCLIConfig is the configuration for the Pokedex CLI application.
var PokedexCLIConfig = s.PokedexCLIConfig{
	StartingMapsAPIEndpoint: PokedexAPIDomain + PokemonMapsURI + "?" + offsetParam + "=" + Offset + "&" + limitParam + "=" + ListSize,
	Prompt:                  PokedexPrompt,
}
