package minecraft

import (
	mc "core-system/logic/minecraft"
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
func StartMinecraft(w http.ResponseWriter, r *http.Request) {
	mc.StartMinecraft(w, r)

	jsonResponse, jsonError := json.Marshal(structs.ApiError{
		Message: "Minecraft started",
		Code:    200,
	})

	if jsonError != nil {
		s.Logger.Println("Unable to encode JSON")
	}

	w.Write(jsonResponse)
}
