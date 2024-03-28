package repl

import (
	"fmt"
	"testing"

	"github.com/KerickHowlett/pokedexcli/internal/tests/utils"
)

func TestWithCommandExecutor(t *testing.T) {
	fmt.Println("WithFields should set the Fields field of the Sanitizer struct")
	repl := NewREPL(WithCommandExecutor(commandExecutorMock))

	utils.ExpectSameEntity(t, repl.CommandExecutor, commandExecutorMock, "CommandExecutor")
}
