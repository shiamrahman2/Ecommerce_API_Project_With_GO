package user

import (
	"ecomerce/config"
	"ecomerce/database"
	"ecomerce/util"
	"encoding/json"
	"net/http"
)

type LogUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func(h* Handler) Login(w http.ResponseWriter, r *http.Request) {
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
	cnf:=config.GetConfig()
	accessToken,err:=util.CreateJWT(cnf.JwtSecretKey,util.PayLoad{
		Sub:usr.ID,
		FirstName:usr.FirstName,
		LastName:usr.LastName,
		Email: usr.Email,
		IsShopOwner:usr.IsShopOwner,
	})
	if err!=nil{
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
		return;
	}
	util.SendData(w,accessToken, http.StatusCreated)
}
