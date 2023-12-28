package commands

import "os"

const OK_EXIT_STATUS = 0

// Exits the REPL (Read-Eval-Print Loop) by calling os.Exit with an "OK" exit
// status of zero (0).
func ExitREPL() error {
	os.Exit(OK_EXIT_STATUS)

	return nil
}
