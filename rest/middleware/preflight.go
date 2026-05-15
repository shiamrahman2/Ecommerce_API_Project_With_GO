package middleware

import "net/http"

func PreFlight(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r) //IF method isn't OPTIONS then route continue it's next step
	})
}
