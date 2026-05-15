package handlers

import (
	"ecomerce/database"
	"ecomerce/util"
	"encoding/json"
	"net/http"
)

type LogUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var reqUser LogUser
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqUser)
	if err != nil {
		println(err)
		http.Error(w, "Invalid User Information", http.StatusBadRequest)
		return
	}
	usr:=database.Find(reqUser.Email,reqUser.Password)
	if usr==nil{
		http.Error(w,"Invalid Credential",http.StatusBadRequest)
		return
	}
	util.SendData(w,usr, http.StatusCreated)
}
