package minecraft

import (
	"github.com/gorilla/mux"
)

func PrepareRouter(router *mux.Router) {
	router.HandleFunc("/install-minecraft", InstallMinecraft).Methods("GET")

	router.HandleFunc("/start-minecraft", StartMinecraft).Methods("GET")
}
