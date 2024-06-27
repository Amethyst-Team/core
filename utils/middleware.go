package utils

import (
	l "core-system/utils/system"
	"net/http"
)

func LogMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		l.Logger.Printf("%s - %s (%s)", r.Method, r.URL.Path, r.RemoteAddr)

		w.Header().Set("Content-Type", "application/json")

		// compare the return-value to the authMW
		next.ServeHTTP(w, r)
	})
}
