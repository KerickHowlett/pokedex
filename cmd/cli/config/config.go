package config

// @TODO: Implement means to pull certain property values from dotenv files
//        based on the environment the application is running in.
//        e.g., dev, qa, prod, etc.

const (
	ListSize         = "20"
	Offset           = "0"
	PokedexAPIDomain = "https://pokeapi.co/api/v2"
	PokedexPrompt    = "pokedex > "
	PokemonMapsURI   = "/location"

	offsetParam = "offset"
	limitParam  = "limit"
)

var (
	StartingMapsAPIEndpoint = PokedexAPIDomain + PokemonMapsURI + "?" + offsetParam + "=" + Offset + "&" + limitParam + "=" + ListSize
)
