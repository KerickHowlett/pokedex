package main

import s "shell"

// @TODO: Implement means to pull certain property values from dotenv files
//        based on the environment the application is running in.
//        e.g., dev, qa, prod, etc.

// PokedexCLIConfig is a struct that holds the configuration for the Pokedex CLI.
//
// Configurations:
//   - limitParam: The query parameter for the number of items to display in a list.
//   - listSize: The number of items to display in a list.
//   - offset: The starting index of the list.
//   - offsetParam: The query parameter for the starting index of the list.
//   - pokedexAPIDomain: The domain of the Pokedex API.
//   - pokedexPrompt: The text displayed as the Pokedex prompt.
//   - pokemonMapsURI: The URI for the Pokemon Maps service.
//   - prompt: The text displayed as the command prompt.
//   - pokemonLocationAreaURI: The URI for the Pokemon Location Area service.
const (
	// The query parameter for the starting index of the list.
	limitParam = "limit"
	// The number of items to display in a list.
	listSize = "20"
	// The starting index of the list.
	offset = "0"
	// The query parameter for the number of items to display in a list.
	offsetParam = "offset"
	// The domain of the Pokedex API.
	pokedexAPIDomain = "https://pokeapi.co/api/v2"
	// The text displayed as the Pokedex prompt.
	pokedexPrompt = "pokedex > "
	// The URI for the Pokemon service.
	pokemonURI = "/pokemon"
	// The URI for the Pokemon Maps service.
	pokemonMapsURI = "/location"
	// The URI for the Pokemon Location Area service.
	pokemonLocationAreaURI = "/location-area"
)

// PokedexCLIConfig is the configuration for the Pokedex CLI application.
var PokedexCLIConfig = s.PokedexCLIConfig{
	LocalAreaAPIEndpoint:    pokedexAPIDomain + pokemonLocationAreaURI,
	PokemonAPIEndpoint:      pokedexAPIDomain + pokemonURI,
	Prompt:                  pokedexPrompt,
	StartingMapsAPIEndpoint: pokedexAPIDomain + pokemonMapsURI + "?" + offsetParam + "=" + offset + "&" + limitParam + "=" + listSize,
}
