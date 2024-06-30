package java

import (
	s "core-system/utils/system"
	"log"
)

// IsInstalled checks if Java is installed on the system.
// It uses the Exec function from the utils/system package to execute the 'java -version' command.
// If the command execution returns an error, it means Java is not installed, and the function returns that error.
// If the command execution is successful, it means Java is installed, and the function returns nil.
func IsInstalled() error {
	log.Printf("Checking if Java is installed...")

	_, err := s.Exec("java -version")

	if err != nil {
		return err
	} else {
		return nil
	}
}
