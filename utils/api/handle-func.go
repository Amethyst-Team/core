package api_utils

import (
	"net/http"
)

func HandleFunc(route string, fun func(w http.ResponseWriter, r *http.Request)) {
	//main.Router.HandleFunc(route, fun)
}
