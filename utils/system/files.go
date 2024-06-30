package system

import (
	"os"
	"path"
)

// CreateFolder creates a new folder at the specified path.
// If the folder already exists, it does nothing.
//
// Parameters:
// folderPath (string): The path of the folder to be created.
//
// Returns:
// This function does not return any value.
//
// Example:
// CreateFolder("/home/user/documents/new_folder")
func CreateFolder(folderPath string) error {
	// Check if the folder already exists.
	fullPath := path.Join(Pwd, folderPath)

	err := os.MkdirAll(fullPath, 0755)

	return err
}

// MoveFile moves a file from its current location to a specified folder.
// If the folder does not exist, it will be created.
//
// Parameters:
// filePath (string): The path of the file to be moved.
// folderPath (string): The path of the destination folder.
//
// Returns:
// This function does not return any value.
//
// Example:
// MoveFile("/home/user/documents/file.txt", "/home/user/downloads")
func MoveFile(filePath string, folderPath string) error {
	// Check if the folder already exists.
	fullPath := path.Join(Pwd, folderPath)

	err := os.MkdirAll(fullPath, 0755)

	if err != nil {
		return err
	}

	err = os.Rename(filePath, path.Join(fullPath, path.Base(filePath)))

	return err
}

func CreateFile(filePath string, content string) error {
	err := os.WriteFile(filePath, []byte(content), 0644)

	return err
}
