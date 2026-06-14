package user

import (
	"ecomerce/config"
)

type Handler struct {
	cnf *config.Config
	svc Service
}

func NewHandler(cnf *config.Config, svc Service) *Handler { // NewHandler Func return an empty Handler Obj. Pointer
	return &Handler{
		cnf: cnf,
		svc: svc,
	}
}
