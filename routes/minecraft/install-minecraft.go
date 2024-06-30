package minecraft

import (
	mc "core-system/logic/minecraft"
	api_utils "core-system/utils/api"
	"net/http"
)

// InstallMinecraft is a function that handles the installation of Minecraft server.
// It takes an http.ResponseWriter and an http.Request as parameters.
// The function calls the InstallMinecraft function from the "core-system/logic/minecraft" package.
// If the return value of InstallMinecraft has a Code field that is not equal to 200,
// the function writes a HTTP status code 400 (Bad Request) to the response writer and sends the return value using the SendResponse function from the "core-system/utils/api" package.
// If the return value of InstallMinecraft has a Code field that is equal to 200,
// the function writes a HTTP status code 200 (OK) to the response writer and sends the return value using the SendResponse function from the "core-system/utils/api" package.
func InstallMinecraft(w http.ResponseWriter, r *http.Request) {
	ret := mc.InstallMinecraft()

	if ret.Code != 200 {
		w.WriteHeader(http.StatusBadRequest)
		api_utils.SendResponse(ret, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	api_utils.SendResponse(ret, w)
}
