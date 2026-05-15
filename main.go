package main

import (
	"ecomerce/util"
	"fmt"
)

//"ecomerce/cmd"

func main() {
	//	cmd.Serve()
	str, err := util.CreateJWT("12345", util.PayLoad{
		Sub:         56,
		FirstName:   "Shiam Hosen",
		LastName:    "Monna",
		Email:       "shiamhosenmona@gmail.com",
		IsShopOwner: false,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(str)
}
