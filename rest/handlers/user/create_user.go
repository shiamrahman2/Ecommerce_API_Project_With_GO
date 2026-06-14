package user

import (
	"ecomerce/domain"
	"ecomerce/util"
	"encoding/json"
	"net/http"
)

type CreatedUser struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser CreatedUser
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)
	if err != nil {
		println(err)
		util.SendError(w, "Invalid Request Data", http.StatusBadRequest)
		return
	}
	createdNewUser, err := h.svc.Create(domain.User{
		FirstName:   newUser.FirstName,
		LastName:    newUser.LastName,
		Email:       newUser.Email,
		Password:    newUser.Password,
		IsShopOwner: newUser.IsShopOwner,
	})
	if err != nil {
		util.SendError(w, "Internal Server Error ", http.StatusInternalServerError)
		return
	}
	util.SendData(w, createdNewUser, http.StatusCreated)
}
