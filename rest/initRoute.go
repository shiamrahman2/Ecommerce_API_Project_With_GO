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
			middleware.PracticeMiddleWare,
		),
	)

	mux.Handle("POST /Products",
		manager.With(
			http.HandlerFunc(handlers.CreateProduct),
				middleware.PracticeMiddleWare,
		),
	)

	mux.Handle("GET /Products/{id}",
		manager.With(
			http.HandlerFunc(handlers.GetProduct),
				middleware.PracticeMiddleWare,
		),
	)

	mux.Handle("PUT /Products/{id}",manager.With(
		http.HandlerFunc(handlers.UpdateProduct),
	  ),
	)
	mux.Handle("DELETE /Products/{id}",manager.With(
		http.HandlerFunc(handlers.DeleteProduct),
	),
	 )
	mux.Handle("POST /user",manager.With(
		 http.HandlerFunc(handlers.CreateUser),
	 ),
	)
	mux.Handle("POST /user/login",manager.With(
		 http.HandlerFunc(handlers.Login),
	 ),
	)
	
}
