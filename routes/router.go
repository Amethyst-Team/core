package router

import (
	"core-system/routes/java"
	"core-system/routes/minecraft"
	"net/http"

	"github.com/gorilla/mux"
)

func PrepareRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router = router.PathPrefix("/api").Subrouter()
	router.HandleFunc("/health", apiHealth)

	minecraftRouter := router.PathPrefix("/minecraft").Subrouter()
	minecraft.PrepareRouter(minecraftRouter)

	javaRouter := router.PathPrefix("/java").Subrouter()
	java.PrepareRouter(javaRouter)

	return router
}

func apiHealth(http.ResponseWriter, *http.Request) {}
