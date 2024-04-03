package repl

import (
	"fmt"
	"testing"

	m "testtools/mocks/parser"
	"testtools/utils"
)

func TestWithParseInput(t *testing.T) {
	fmt.Println("WithParseInput should set the ParseInput field of the REPL struct")

	parser := m.MockParser
	repl := setupREPLOptionTest(WithParseInput(parser))

	utils.ExpectSameEntity(t, repl.ParseInput, parser, "ParseInput")
}
