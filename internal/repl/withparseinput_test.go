package repl

import (
	"fmt"
	"testing"

	m "github.com/KerickHowlett/pokedexcli/tests/mocks"
	"github.com/KerickHowlett/pokedexcli/tests/utils"
)

func TestWithParseInput(t *testing.T) {
	fmt.Println("WithParseInput should set the ParseInput field of the REPL struct")

	parser := m.ParserMock
	repl := setupREPLOptionTest(WithParseInput(parser))

	utils.ExpectSameEntity(t, repl.ParseInput, parser, "ParseInput")
}
