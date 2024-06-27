package router

import (
	"core-system/routes/minecraft"
	"net/http"

	"github.com/gorilla/mux"
)

func PrepareRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	// router.HandleFunc("/api/health", apiHealth)

	minecraftRouter := router.PathPrefix("/minecraft").Subrouter()
	minecraft.PrepareRouter(minecraftRouter)

	return router
}

func apiHealth(http.ResponseWriter, *http.Request) {}
