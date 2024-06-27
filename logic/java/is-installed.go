package java

import (
	s "core-system/utils/system"
)

func IsInstalled() error {
	s.Logger.Printf("Checking if Java is installed...")

	_, err := s.Exec("java -version")
	if err != nil {
		s.Logger.Printf("Java is not installed.")
		return err
	} else {
		s.Logger.Printf("Java is installed.")
		return nil
	}
}
