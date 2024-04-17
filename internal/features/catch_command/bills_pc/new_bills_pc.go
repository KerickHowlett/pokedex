package bills_pc

// NewBillsPC creates a new BillsPC state with the provided options set in it.
//
// Parameters:
//   - options: a list of BillsPCOption functions that will be applied to the new BillsPC state.
//
// Returns:
//   - a new BillsPC state with the provided options set in it.
//
// Example:
//
//	billsPC := NewBillsPC()
func NewBillsPC(options ...BillsPCOption) *BillsPC {
	billsPC := &BillsPC{
		caughtPokemon: &CaughtPokemon{},
		withdrawError: "you have not caught this pokemon",
	}
	for _, option := range options {
		option(billsPC)
	}
	return billsPC
}
