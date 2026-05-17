package user

import (
	"ecomerce/database"
	"ecomerce/util"
	"encoding/json"
	"net/http"
)

func(h* Handler) CreateUser(w http.ResponseWriter,r * http.Request) {
     var newUser database.User
     decoder:=json.NewDecoder(r.Body)
	 err:=decoder.Decode(&newUser)
	 if err!=nil{
		println(err)
		http.Error(w,"Invalid Request Data",http.StatusBadRequest)
		return
	 }
	 createdNewUser:=newUser.Store()
	 util.SendData(w,createdNewUser,http.StatusCreated)
}