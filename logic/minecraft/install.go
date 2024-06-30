package minecraft

import (
	"core-system/structs"
	s "core-system/utils/system"
)

const mcUrl = "http://piston-data.mojang.com/v1/objects/450698d1863ab5180c25d7c804ef0fe6369dd1ba/server.jar"
const mcServerPath = "server-data\\minecraft\\server"

// InstallMinecraft function is responsible for downloading the Minecraft server jar file,
// creating a server folder if it doesn't exist, and moving the downloaded file to the server folder.
// It returns an ApiError struct if any error occurs during the process.
func InstallMinecraft() structs.ApiError {
	// Download the Minecraft server jar file from the provided URL.
	filePath, err := s.DownloadFile("minecraft", mcUrl, false)

	// If an error occurred during the download, return an ApiError with the error message and a 400 status code.
	if err != nil {
		return structs.ApiError{
			Message: err.Error(),
			Code:    400,
		}
	}

	// Create a server folder if it doesn't exist.
	s.CreateFolder(mcServerPath)

	// Move the downloaded file to the server folder.
	s.MoveFile(filePath, mcServerPath)

	// Return an ApiError with a success message and a 200 status code.
	return structs.ApiError{
		Message: "Successfully installed Minecraft",
		Code:    200,
	}
}
