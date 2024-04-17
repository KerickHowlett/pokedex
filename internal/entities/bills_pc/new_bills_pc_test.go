package bills_pc

import "testing"

func TestNewBillsPC(t *testing.T) {
	t.Run("should create a new non-nil BillsPC instance.", func(t *testing.T) {
		if billsPC := NewBillsPC(); billsPC == nil {
			t.Errorf("NewBillsPC() = nil; want non-nil BillsPC instance.")
		}
	})
}
