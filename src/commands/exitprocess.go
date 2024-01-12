package commands

import "os"

// Exits the REPL (Read-Eval-Print Loop) by calling os.Exit with an "OK" exit
// status of zero (0).
func ExitProcess() error {
	const OKAY_STATUS = 0
	os.Exit(OKAY_STATUS)

	return nil
}
