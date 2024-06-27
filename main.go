package main

import (
	"log"
	"net/http"

	router "core-system/routes"
)

func main() {
	router := router.PrepareRouter()

	Address := Options.LISTEN_ADDRESS + ":" + Options.LISTEN_PORT

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	srv := &http.Server{
		Handler: router,
		Addr:    Address,
	}

	log.Printf("Listening on %s\n", Address)
	log.Fatal(srv.ListenAndServe())
}
