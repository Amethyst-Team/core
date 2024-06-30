package minecraft

import (
	mc "core-system/logic/minecraft"
	api_utils "core-system/utils/api"
	"net/http"
)

// StartMinecraft starts the Minecraft server.
//
// This function takes two parameters:
// - w: http.ResponseWriter to write the HTTP response.
// - r: *http.Request to handle the HTTP request.
//
// The function returns nothing.
//
// It calls the StartMinecraft function from the "core-system/logic/minecraft" package.
// If the returned code from the StartMinecraft function is not 200, it writes a HTTP status code 400 (Bad Request) and sends the response using the SendResponse function from the "core-system/utils/api" package.
// If the returned code from the StartMinecraft function is 200, it writes a HTTP status code 200 (OK) and sends the response using the SendResponse function from the "core-system/utils/api" package.
func StartMinecraft(w http.ResponseWriter, r *http.Request) {
	ret := mc.StartMinecraft()

	if ret.Code != 200 {
		w.WriteHeader(http.StatusBadRequest)
		api_utils.SendResponse(ret, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	api_utils.SendResponse(ret, w)
}
