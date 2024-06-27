package system

import (
	"runtime"

	"github.com/abdfnx/gosh"
)

// Exec executes a shell command and returns its output.
// If the command execution fails, it logs the error and returns the error message.
//
// Parameters:
// cmd (string): The shell command to be executed.
//
// Returns:
// (string, error): A tuple containing the command output and an error if any.
// If the command execution is successful, the error will be nil.
func Exec(cmd string) (string, error) {
	var err error
	var out, errout string

	// Check the operating system and execute the command accordingly.
	if runtime.GOOS == "windows" {
		err, out, errout = gosh.PowershellOutput(cmd)
	} else {
		err, out, errout = gosh.ShellOutput(cmd)
	}

	// If an error occurred during command execution, return it.
	if err != nil {
		return errout, err
	}

	// If the command execution is successful, return the output.
	return out, nil
}
