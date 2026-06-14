package user

import (
	"ecomerce/util"
	"encoding/json"
	"net/http"
)

type LogUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var reqUser LogUser
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqUser)
	if err != nil {
		println(err)
		http.Error(w, "Invalid User Information", http.StatusBadRequest)
		return
	}
	usr, err := h.svc.Find(reqUser.Email, reqUser.Password)
	if err != nil {
		util.SendError(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if usr == nil {
		util.SendError(w, "Invalid Credential", http.StatusBadRequest)
		return
	}
	accessToken, err := util.CreateJWT(h.cnf.JwtSecretKey, util.PayLoad{
		Sub:         usr.ID,
		FirstName:   usr.FirstName,
		LastName:    usr.LastName,
		Email:       usr.Email,
		IsShopOwner: usr.IsShopOwner,
	})
	if err != nil {
		util.SendError(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	util.SendData(w, accessToken, http.StatusCreated)
}
