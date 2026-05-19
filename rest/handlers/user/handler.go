package user

import (
	"ecomerce/config"
	"ecomerce/repo"
)

type Handler struct {
	cnf      *config.Config
	userRepo repo.UserRepo
}

func NewHandler(cnf *config.Config, userRepo repo.UserRepo) *Handler { // NewHandler Func return an empty Handler Obj. Pointer
	return &Handler{
		cnf:      cnf,
		userRepo: userRepo,
	}
}
