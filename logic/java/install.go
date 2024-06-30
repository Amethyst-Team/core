package java

import (
	s "core-system/utils/system"
	"strings"
)

const javaUrl = "https://download.oracle.com/java/22/latest/jdk-22_windows-x64_bin.exe"

// InstallJava downloads and installs the latest version of Java from Oracle's website.
// It uses the provided utility functions from the "core-system/utils/system" package.
// The function first downloads the Java installer file to a temporary location.
// If the download is successful, it executes the installer with the "/s" flag to perform a silent installation.
// If any error occurs during the download or installation process, the function panics with the error.
func InstallJava() error {
	filePath, err := s.DownloadFile("java", javaUrl, false)

	if err != nil {
		panic(err)
	}

	_, err = s.Exec(strings.Join([]string{filePath, "/s"}, " "))

	if err != nil {
		return err
	}

	err = s.RestartSelf()

	return err
}
