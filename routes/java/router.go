package java

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"core-system/logic/java"
	structs "core-system/structs"
	s "core-system/utils/system"
	//mc "core-system/logic/minecraft"
)

func PrepareRouter(router *mux.Router) {
	router.HandleFunc("/installjava", installJava)
	router.HandleFunc("/isjavainstalled", isInstalled)
}

// installJava handles the installation of Java.
// It checks if Java is already installed. If not, it attempts to install Java.
// It returns a JSON response with appropriate message and HTTP status code.
// If Java is already installed, it returns a 400 status code with a message indicating that Java is already installed.
// If the installation fails, it returns a 400 status code with a message indicating the failure.
// If the installation is successful, it returns a 200 status code with a message indicating the success.
func installJava(w http.ResponseWriter, r *http.Request) {
	err := java.IsInstalled()

	// java is installed
	if err == nil {
		jsonResponse, jsonError := json.Marshal(structs.ApiError{
			Message: "Java is already installed",
			Code:    200,
		})

		if jsonError != nil {
			s.Logger.Println("Unable to encode JSON")
		}

		w.Write(jsonResponse)
		return
	}

	err = java.InstallJava()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		jsonResponse, jsonError := json.Marshal(structs.ApiError{
			Message: "Failed to install Java",
			Code:    400,
		})

		if jsonError != nil {
			s.Logger.Println("Unable to encode JSON")
		}

		w.Write(jsonResponse)
		return
	}

	jsonResponse, jsonError := json.Marshal(structs.ApiError{
		Message: "Successfully installed Java",
		Code:    200,
	})

	if jsonError != nil {
		s.Logger.Println("Unable to encode JSON")
	}

	w.Write(jsonResponse)
}

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
			s.Logger.Println("Unable to encode JSON")
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
		s.Logger.Println("Unable to encode JSON")
	}

	// Write JSON response to the response writer
	w.Write(jsonResponse)
}
