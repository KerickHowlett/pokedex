package repl

import (
	"fmt"
	"testing"

	m "testtools/mocks/scanner"
	u "testtools/utils"
)

func TestWithScanner(t *testing.T) {
	fmt.Println("WithScanner should set the Scanner field of the REPL struct")

	scanner := m.NewMockScanner()
	repl := setupREPLOptionTest(WithScanner(scanner))

	u.ExpectSameEntity(t, repl.Scanner, scanner, "Scanner")
}
