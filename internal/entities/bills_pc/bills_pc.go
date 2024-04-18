package bills_pc

import (
	"errors"
	p "pokemon"
)

// CaughtPokemon is a map of all the caught Pokemon indexed by their respective names.
type CaughtPokemon map[string]*p.Pokemon

// BillsPC is a struct that represents the state of the user's BillsPC.
//
// Fields:
//   - caughtPokemon: a map of all the caught Pokemon indexed by their respective names.
//   - withdrawError: the error message that will be returned when a Pokemon is not found in the BillsPC.
//
// Methods:
//   - Deposit: adds a Pokemon to the BillsPC if it has not already been caught.
//   - Inspect: returns the Pokemon with the provided name if it has been caught.
//   - GetCaughtPokemon: returns a map of all the caught Pokemon indexed by their respective names.
//   - Withdraw: removes a Pokemon from the BillsPC if it has been caught.
type BillsPC struct {
	// caughtPokemon is a map of all the caught Pokemon indexed by their respective names.
	caughtPokemon *CaughtPokemon
	// withdrawError is the error message that will be returned when a Pokemon is not found in the BillsPC.
	withdrawError string
}

// Deposit adds a Pokemon to the BillsPC if it has not already been caught.
//
// Parameters:
//   - pokemon: the Pokemon to be added to the BillsPC.
//
// Example usage:
//
//	billsPC := NewBillsPC()
//	billsPC.Deposit(pokemon)
func (b BillsPC) Deposit(pokemon *p.Pokemon) {
	if _, alreadyCaught := b.Inspect(pokemon.Name); !alreadyCaught {
		(*b.caughtPokemon)[pokemon.Name] = pokemon
	}
}

// Inspect returns the Pokemon with the provided name if it has been caught.
//
// Parameters:
//   - pokemonName: the name of the Pokemon to be inspected.
//
// Returns:
//   - the Pokemon with the provided name if it has been caught.
//   - a boolean indicating whether the Pokemon has been caught.
//
// Example usage:
//
//	billsPC := NewBillsPC()
//	pokemon, found := billsPC.Inspect("Pikachu")
func (b BillsPC) Inspect(pokemonName string) (pokemon *p.Pokemon, found bool) {
	pokemon, found = (*b.caughtPokemon)[pokemonName]
	return pokemon, found
}

// GetCaughtPokemon returns a map of all the caught Pokemon indexed by their respective names.
//
// Returns:
//   - a map of all the caught Pokemon indexed by their respective names.
//
// Example usage:
//
//	billsPC := NewBillsPC()
//	caughtPokemon := billsPC.GetCaughtPokemon()
func (b BillsPC) GetCaughtPokemon() CaughtPokemon {
	return *b.caughtPokemon
}

// TotalCaughtPokemon returns the total number of Pokemon caught by the user.
//
// Returns:
//   - the total number of Pokemon caught by the user.
//
// Example usage:
//
//	billsPC := NewBillsPC()
//	totalCaughtPokemon := billsPC.TotalCaughtPokemon()
func (b BillsPC) TotalCaughtPokemon() int {
	return len(*b.caughtPokemon)
}

// Withdraw removes a Pokemon from the BillsPC if it has been caught.
//
// Parameters:
//   - pokemonName: the name of the Pokemon to be removed from the BillsPC.
//
// Returns:
//   - the Pokemon that was removed from the BillsPC.
//   - an error if the Pokemon was not found in the BillsPC.
//
// Example usage:
//
//	billsPC := NewBillsPC()
//	pokemon, err := billsPC.Withdraw("Pikachu")
func (b *BillsPC) Withdraw(pokemonName string) (pokemon *p.Pokemon, err error) {
	pokemon, found := b.Inspect(pokemonName)
	if !found {
		return nil, errors.New(b.withdrawError)
	}

	delete(*b.caughtPokemon, pokemonName)

	return pokemon, nil
}
