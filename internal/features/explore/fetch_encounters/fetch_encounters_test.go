package fetch_encounters

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	la "entities/location_area"
	p "entities/pokemon"
	pe "entities/pokemon_encounter"
	qf "query/fetch"
	f "test_tools/fixtures"
	"test_tools/utils"
)

func TestFetchEncounters(t *testing.T) {
	setup := func(responseType string) (state *la.LocationArea, mockedFetch qf.QueryFetchFunc[la.LocationArea], stdout *utils.PrintStorage) {
		state, stdout = la.NewLocationArea(), utils.NewPrintStorage()

		fetchedMock := func(responseType string) qf.QueryFetchFunc[la.LocationArea] {
			if responseType == "error" {
				return func(url string, query *la.LocationArea, ttl ...time.Duration) error {
					return fmt.Errorf("error with response: %s", http.StatusText(http.StatusInternalServerError))
				}
			}
			if responseType == "no-maps" {
				return func(url string, query *la.LocationArea, ttl ...time.Duration) error {
					query.Encounters = []pe.PokemonEncounter{}
					return nil
				}
			}
			if responseType == "success" {
				return func(url string, query *la.LocationArea, ttl ...time.Duration) error {
					pokemon := p.NewPokemon(p.WithName(f.PokemonName))
					encounter := pe.NewPokemonEncounter(pe.WithPokemon(pokemon))
					query.Encounters = append(query.Encounters, *encounter)
					return nil
				}
			}
			panic(fmt.Sprintf("Unexpected responseType used for test setup: %s", responseType))
		}(responseType)

		return state, fetchedMock, stdout
	}

	expectErrorMessage := func(err error, expectedErrorMessage string) {
		if err == nil {
			t.Error("Expected an error, but got nil.")
		}
		if actualErrorMessage := err.Error(); actualErrorMessage != expectedErrorMessage {
			t.Errorf("Expected error to be %q, but got %q.", expectedErrorMessage, actualErrorMessage)
		}
	}

	t.Run("should return an error if the state is nil.", func(t *testing.T) {
		t.Parallel()
		_, mockedFetch, _ := setup("success")
		err := FetchEncounters(f.APIEndpoint, nil, mockedFetch)
		expectErrorMessage(err, "a query state is required to fetch locations")
	})

	t.Run("should return an error if the query fetch fails.", func(t *testing.T) {
		t.Parallel()
		state, mockedFetch, _ := setup("error")
		expectedErrorMessage := fmt.Sprintf("error with response: %s", http.StatusText(http.StatusInternalServerError))
		err := FetchEncounters(f.APIEndpoint, state, mockedFetch)
		expectErrorMessage(err, expectedErrorMessage)
	})

	t.Run("should return an error if no locations are found.", func(t *testing.T) {
		t.Parallel()
		state, mockedFetch, _ := setup("no-maps")
		err := FetchEncounters(f.APIEndpoint, state, mockedFetch)
		expectErrorMessage(err, "no pokemon were found in the area")
	})

	t.Run("should print all map location names stored in query state.", func(t *testing.T) {
		t.Parallel()
		state, mockedFetch, stdout := setup("success")

		output, err := stdout.CaptureWithError(func() error {
			return FetchEncounters(f.APIEndpoint, state, mockedFetch)
		})

		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if expectedOutput := fmt.Sprintf("Found Pokemon:\n  - %s\n", f.PokemonName); output != expectedOutput {
			t.Errorf("Expected output to be %q, but got %q.", expectedOutput, output)
		}
	})
}
