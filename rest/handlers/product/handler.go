package product

import (
	"ecomerce/rest/middleware"
)

type Handler struct {
	middleware *middleware.MiddleWare
	svc Service
}

func NewHandler(middleware *middleware.MiddleWare,
	 svc Service,
	   ) *Handler {
	  return &Handler{
		middleware:middleware,
		svc:svc,
	}
}