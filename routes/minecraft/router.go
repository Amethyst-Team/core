package minecraft

import (
	"net/http"

	"github.com/gorilla/mux"
)

func PrepareRouter(router *mux.Router) {
	router.HandleFunc("/hello", func(http.ResponseWriter, *http.Request) {})
}
