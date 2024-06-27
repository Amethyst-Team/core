package utils

import (
	"log"
	"runtime"

	"github.com/abdfnx/gosh"
)

func Exec(cmd string) (string, error) {
	var err error
	var out, errout string

	if runtime.GOOS == "windows" {
		err, out, errout = gosh.PowershellOutput(cmd)
	} else {
		err, out, errout = gosh.ShellOutput(cmd)
	}

	if err != nil {
		log.Printf("Error while executing: %v\n", err)

		return errout, err
	}

	return out, nil
}
