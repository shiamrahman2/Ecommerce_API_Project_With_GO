package rest

import (
	"ecomerce/rest/handlers"
	"ecomerce/rest/middleware"
	"net/http"
)

func InitRoute(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("GET /Products",
		manager.With(
			http.HandlerFunc(handlers.GetProducts),
		),
	)

	mux.Handle("POST /Products",
		manager.With(
			http.HandlerFunc(handlers.CreateProduct),
			middleware.AuthenticationJWT,
		),
	)

	mux.Handle("GET /Products/{id}",
		manager.With(
			http.HandlerFunc(handlers.GetProduct),
		),
	)

	mux.Handle("PUT /Products/{id}", manager.With(
		http.HandlerFunc(handlers.UpdateProduct),
		middleware.AuthenticationJWT,
	),
	)
	mux.Handle("DELETE /Products/{id}", manager.With(
		http.HandlerFunc(handlers.DeleteProduct),
		middleware.AuthenticationJWT,
	),
	)
	mux.Handle("POST /user", manager.With(
		http.HandlerFunc(handlers.CreateUser),
	),
	)
	mux.Handle("POST /user/login", manager.With(
		http.HandlerFunc(handlers.Login),
	),
	)

}
