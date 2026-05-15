package main

import (
	"crypto/hmac"
	"crypto/sha256"
	//"encoding/base64"
	"fmt"
)

// import (
// 	"ecomerce/cmd"
// )

func main() {
	//cmd.Serve()
	
	secret:=[]byte("my-secret")
	message:=[]byte("Hello World")
    
	hash:=hmac.New(sha256.New,secret)
	hash.Write(message)
	hashFinal:=hash.Sum(nil)
	fmt.Println(hashFinal)
}
