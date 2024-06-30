package system

import (
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

var downloads = "downloads"

var downloadDir string
var Pwd string

// init function initializes the working directory for downloads.
// It sets the working directory to the current directory and creates a "downloads" folder if it doesn't exist.
func init() {
	// Get the current working directory.
	dir, err := os.Getwd()

	// If an error occurs while getting the working directory, log the error and exit the program.
	if err != nil {
		ErrorLog.Printf("Error while getting working directory: %v\n", err)
		os.Exit(1)
	}

	// Attempt to create the "downloads" folder in the working directory.
	os.Mkdir(downloads, os.ModePerm)

	// Set the working directory for downloads to the "downloads" folder within the current working directory.
	downloadDir = path.Join(dir, downloads)
	Pwd = dir
}

// DownloadFile downloads a file from the given URL to the specified directory.
// If the directory does not exist, it will be created.
// If the file already exists and overwrite is set to true, the existing file will be overwritten.
// The function returns the path of the downloaded file and any encountered errors.
func DownloadFile(directory string, url string, overwrite bool) (filePath string, err error) {
	// Construct the download directory path
	downloadDir := path.Join(downloadDir, directory)

	// Create the download directory if it does not exist
	err = os.MkdirAll(downloadDir, os.ModePerm)

	// Extract the file name from the URL
	urlSplit := strings.Split(url, "/")
	fileName := urlSplit[len(urlSplit)-1]

	// Construct the path to save the downloaded file
	pathToSave := path.Join(downloadDir, "\\"+fileName)

	// Remove the existing file if overwrite is set to true
	if overwrite {
		os.Remove(pathToSave)
	}

	// Create a new file for writing the downloaded content
	out, err := os.Create(pathToSave)
	if err != nil {
		return "", err
	}
	defer out.Close()

	// Send an HTTP GET request to the URL
	resp, err := http.Get(url)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Write the response body to the file
	err = os.WriteFile(pathToSave, body, 0644)
	if err != nil {
		return "", err
	}

	// Check the HTTP status code
	if resp.StatusCode != http.StatusOK {
		ErrorLog.Printf("failed to download file: %s", resp.Status)
		return "", err
	}

	// Copy the response body to the file (redundant as we already read the body)
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", err
	}
	log.Print(pathToSave)
	// Return the path of the downloaded file
	return pathToSave, nil
}
