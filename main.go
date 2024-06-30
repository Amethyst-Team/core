package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"runtime"

	router "core-system/routes"
	u "core-system/utils"
	walk "core-system/utils/api"
	w "core-system/windows"
)

// main function is the entry point of the application.
// It prepares the router, sets up the server, checks for admin privileges,
// installs Java, and starts listening for incoming connections.
func main() {
	// Check if the application is running with admin privileges.
	// If not, run the application as an elevated process.
	if runtime.GOOS == "windows" {
		if !w.IsAppAdmin() {
			w.RunAppAsElevated()
		}
	}

	elevated := w.IsAppAdmin()

	if !elevated {
		log.Printf("Exiting because we can't escalate permissions to Administrator...")
		os.Exit(1)
	}

	// Prepare the router for handling HTTP requests.
	router := router.PrepareRouter()

	// Combine the listen address and port to form the server address.
	Address := Options.LISTEN_ADDRESS + ":" + Options.LISTEN_PORT

	var dir string

	flag.StringVar(&dir, "dir", "./static/", "the directory to serve files from. Defaults to the current dir")
	flag.Parse()

	// Serve static files from the "./static/" directory.
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(dir))))

	// Add a middleware function to log HTTP requests.
	router.Use(u.LogMW)

	// Create a new HTTP server with the router as the handler.
	srv := &http.Server{
		Handler: router,
		Addr:    Address,
	}

	walk.PrintEndpoints(router)

	// Log the server listening address.
	log.Printf("Listening on %s\n", Address)

	// Start the server and log any errors that occur.
	log.Fatal(srv.ListenAndServe())
}
