package middleware

import "net/http"

func Cors( next http.Handler) http.Handler {
	
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//CORS handle for all request globally
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST,GET,PATCH,DELETE,OPTIONS,PUT")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r) //IF method isn't OPTIONS then route continue it's next step
	})
}
