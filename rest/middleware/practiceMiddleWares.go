package middleware

import (
	"log"
	"net/http"
)

func PracticeMiddleWare(next http.Handler) http.Handler {
   return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	   log.Println("Ami Practice MiddleWares")
        next.ServeHTTP(w,r)
   })
}