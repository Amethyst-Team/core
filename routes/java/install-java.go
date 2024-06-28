package java

import (
	"core-system/logic/java"
	structs "core-system/structs"
	s "core-system/utils/system"
	"encoding/json"
	"net/http"
)

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
