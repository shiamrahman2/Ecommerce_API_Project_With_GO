package user

import (
	"ecomerce/rest/middleware"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("POST /user", manager.With(
		http.HandlerFunc(h.CreateUser),
	),
	)
	mux.Handle("POST /user/login", manager.With(
		http.HandlerFunc(h.Login),
	),
	)

}
