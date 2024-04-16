package fixtures

import "time"

var (
	// APIEndpoint represents the URL of the API endpoint domain. which
	// is used for tests that require an API endpoint.
	APIEndpoint = "https://example.com"

	// FrozenTime represents a fixed time used for tests that require a
	// fixed time.
	// This time is set to April 6, 2024, 12:00:00 UTC.
	FrozenTime = time.Date(2024, time.April, 6, 12, 0, 0, 0, time.UTC)

	// Invalid represents an invalid value used for tests that require an
	// invalid value.
	Invalid = "invalid"

	// EmptyLine represents an empty line used for tests that require an
	// empty line.
	EmptyLine = "\n"

	// PokemonName represents the name of the Pokemon used for tests.
	PokemonName = "Pikachu"

	// StarterTown represents the name of the starting town in the game,
	// which is used for filling out Location.Name for tests.
	StarterTown = "Pallet Town"
)
