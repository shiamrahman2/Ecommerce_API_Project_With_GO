package main

import (
	"crypto/sha256"
	//"encoding/base64"
	"fmt"
)

// import (
// 	"ecomerce/cmd"
// )

func main() {
	//cmd.Serve()
	
	data:=[]byte("Hello World")
	EncryptedData:=sha256.Sum256(data)
	fmt.Println(EncryptedData)

}
