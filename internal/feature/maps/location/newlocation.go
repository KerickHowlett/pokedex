package location

func NewLocation(options ...LocationOption) *Location {
	pokemonMap := &Location{
		Name: "",
	}

	for _, option := range options {
		option(pokemonMap)
	}

	if pokemonMap.Name == "" {
		panic("[ERROR] Location name is required")
	}

	return pokemonMap
}
