package utils

import (
	"net/http"
)

func LogMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		GeneralLogger.Printf("%s - %s (%s)", r.Method, r.URL.Path, r.RemoteAddr)

		// compare the return-value to the authMW
		next.ServeHTTP(w, r)
	})
}
