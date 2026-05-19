package user

import (
	repo "ecomerce/Repo"
	"ecomerce/config"
)

type Handler struct {
     cnf *config.Config
	userRepo repo.UserRepo
}

func NewHandler( cnf *config.Config,userRepo repo.UserRepo) *Handler { // NewHandler Func return an empty Handler Obj. Pointer
	return &Handler{
          cnf:cnf,
		userRepo: userRepo,
	}
}
