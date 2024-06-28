package java

import (
	"github.com/gorilla/mux"
)

// PrepareRouter sets up the routes for the Java installation and checking functionalities.
// It accepts a pointer to a mux.Router and adds two routes:
// - "/install-java" which responds to GET requests and triggers the installJava function.
// - "/is-java-installed" which responds to GET requests and triggers the isInstalled function.
func PrepareRouter(router *mux.Router) {
	// Route for installing Java
	router.HandleFunc("/install-java", installJava).Methods("GET")

	// Route for checking if Java is installed
	router.HandleFunc("/is-java-installed", isInstalled).Methods("GET")
}
