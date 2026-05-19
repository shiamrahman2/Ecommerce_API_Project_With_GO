package product

import (
	"ecomerce/repo"
	"ecomerce/rest/middleware"
)

type Handler struct {
	middleware *middleware.MiddleWare
	productRepo repo.ProductRepo
}

func NewHandler(middleware *middleware.MiddleWare,
	 productRepo repo.ProductRepo,
	   ) *Handler {
	  return &Handler{
		middleware:middleware,
		productRepo:productRepo,
	}
}