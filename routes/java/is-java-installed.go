package java

import (
	"core-system/logic/java"
	structs "core-system/structs"
	"encoding/json"
	"log"
	"net/http"
)

// isInstalled handles the check for Java installation.
// It checks if Java is already installed. If Java is installed, it returns a JSON response with a message indicating that Java is installed and HTTP status code 200.
// If Java is not installed, it returns a JSON response with a message indicating that Java is not installed and HTTP status code 400.
func isInstalled(w http.ResponseWriter, r *http.Request) {
	// Check if Java is installed
	err := java.IsInstalled()

	// If Java is not installed
	if err != nil {
		// Set HTTP status code to 400 (Bad Request)
		w.WriteHeader(http.StatusBadRequest)

		// Create JSON response
		jsonResponse, jsonError := json.Marshal(structs.ApiError{
			Message: "Java is not installed",
			Code:    400,
		})

		// If there is an error while encoding JSON
		if jsonError != nil {
			// Log the error
			log.Println("Unable to encode JSON")
		}

		// Write JSON response to the response writer
		w.Write(jsonResponse)
		return
	}

	// If Java is installed
	// Create JSON response
	jsonResponse, jsonError := json.Marshal(structs.ApiError{
		Message: "Java installed",
		Code:    200,
	})

	// If there is an error while encoding JSON
	if jsonError != nil {
		// Log the error
		log.Println("Unable to encode JSON")
	}

	// Write JSON response to the response writer
	w.Write(jsonResponse)
}
