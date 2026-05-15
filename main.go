package main

import (
	"encoding/base64"
	"fmt"
)

// import (
// 	"ecomerce/cmd"
// )

/*
handleCORS and handlePreflightRequest don't required because i handle CORS and Preflight With global MUX

	func handleCORS(w http.ResponseWriter) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST,GET,PATCH,DELETE,OPTIONS,PUT")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "application/json")
	}

	func handlePreflightRequest(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
	}
*/
func main() {
	//cmd.Serve()
	var s string
	s="abacabed"
	fmt.Println(s)
    byteArray:=[]byte(s)
	fmt.Println(byteArray)
	enc:=base64.URLEncoding
    enc=enc.WithPadding(base64.NoPadding)
	b64Encoding:=enc.EncodeToString(byteArray)
	fmt.Println(b64Encoding)
}
