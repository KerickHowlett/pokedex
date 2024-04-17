package bills_pc

import p "pokemon"

type BillsPCOption func(*BillsPC)

// WithCaughtPokemon is a BillsPCOption adds a map entry to BillsPC.caughtPokemon.
//
// Parameters:
//   - caughtPokemon: the Pokemon to add to the caughtPokemon map.
//
// Returns:
//   - a BillsPCOption function that adds the caughtPokemon to the caughtPokemon map.
//
// Example usage:
//
//	billsPC := NewBillsPC(WithCaughtPokemon(caughtPokemon))
func WithCaughtPokemon(caughtPokemon *p.Pokemon) BillsPCOption {
	return func(b *BillsPC) {
		(*b.caughtPokemon)[caughtPokemon.Name] = caughtPokemon
	}
}

// WithWithdrawError is a BillsPCOption that sets the BillsPC.withdrawError field.
//
// Parameters:
//   - withdrawError: the error message that will be returned when a Pokemon is not found in the BillsPC.
//
// Returns:
//   - a BillsPCOption function that sets the BillsPC.withdrawError field.
//
// Example usage:
//
//	billsPC := NewBillsPC(WithWithdrawError("you have not caught this pokemon"))
func WithWithdrawError(withdrawError string) BillsPCOption {
	return func(b *BillsPC) {
		b.withdrawError = withdrawError
	}
}
