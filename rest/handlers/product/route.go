package product

import (
	"ecomerce/rest/middleware"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("GET /Products",
		manager.With(
			http.HandlerFunc(h.GetProducts),
		),
	)

	mux.Handle("POST /Products",
		manager.With(
			http.HandlerFunc(h.CreateProduct),
			middleware.AuthenticationJWT,
		),
	)

	mux.Handle("GET /Products/{id}",
		manager.With(
			http.HandlerFunc(h.GetProduct),
		),
	)

	mux.Handle("PUT /Products/{id}", manager.With(
		http.HandlerFunc(h.UpdateProduct),
		middleware.AuthenticationJWT,
	),
	)
	mux.Handle("DELETE /Products/{id}", manager.With(
		http.HandlerFunc(h.DeleteProduct),
		middleware.AuthenticationJWT,
	),
	)

}
