package product

import "ecomerce/rest/middleware"

type Handler struct {
	middleware *middleware.MiddleWare
}

func NewHandler(middleware *middleware.MiddleWare) *Handler {
	return &Handler{
		middleware:middleware,
	}
}