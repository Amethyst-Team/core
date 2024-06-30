package api_utils

import (
	"core-system/structs"
	"encoding/json"
	"net/http"
)

func SendResponse(res structs.ApiError, w http.ResponseWriter) {
	// Set the response status code
	w.WriteHeader(res.Code)

	// Set the response header
	SetHeaderJSON(w)

	// Convert the response struct to JSON
	json.NewEncoder(w).Encode(res)
}

func SetHeaderJSON(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}
