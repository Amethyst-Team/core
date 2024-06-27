package main

import (
	"net/http"
	"os"
	"runtime"

	router "core-system/routes"
	u "core-system/utils"
	s "core-system/utils/system"
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
		os.Exit(1)
	} else {
		// Print the result to the console for debugging purposes.
		s.Logger.Printf("Running privileged? %v\n", elevated)
	}

	// Prepare the router for handling HTTP requests.
	router := router.PrepareRouter()

	// Combine the listen address and port to form the server address.
	Address := Options.LISTEN_ADDRESS + ":" + Options.LISTEN_PORT

	// Serve static files from the "./static/" directory.
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	// Add a middleware function to log HTTP requests.
	router.Use(u.LogMW)

	// Create a new HTTP server with the router as the handler.
	srv := &http.Server{
		Handler: router,
		Addr:    Address,
	}

	// Log the server listening address.
	s.Logger.Printf("Listening on %s\n", Address)

	// Start the server and log any errors that occur.
	s.ErrorLog.Fatal(srv.ListenAndServe())
}
