package command

import "os"

const okayStatus = 0

func exitApplication(exitStatus int) error {
	os.Exit(exitStatus)
	return nil
}
