package repl

import (
	"fmt"
	"testing"

	"testtools/utils"
)

func TestWithCommandExecutor(t *testing.T) {
	fmt.Println("should set the Fields field of the Sanitizer struct")
	repl := NewREPL(WithCommandExecutor(commandExecutorMock))

	utils.ExpectSameEntity(t, repl.execute, commandExecutorMock, "execute")
}
