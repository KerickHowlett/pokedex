package commands

import "os"

const ok_exit_status = 0

// Exits the REPL (Read-Eval-Print Loop) by calling os.Exit with an "OK" exit
// status of zero (0).
func ExitREPL() error {
	os.Exit(ok_exit_status)

	return nil
}
