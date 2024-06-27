package main

import (
	"net/http"

	router "core-system/routes"
	u "core-system/utils"
)

func main() {
	router := router.PrepareRouter()

	Address := Options.LISTEN_ADDRESS + ":" + Options.LISTEN_PORT

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	router.Use(u.LogMW)

	srv := &http.Server{
		Handler: router,
		Addr:    Address,
	}

	u.GeneralLogger.Printf("Listening on %s\n", Address)
	u.GeneralLogger.Fatal(srv.ListenAndServe())
}
