package rest

import (
	"ecomerce/config"
	"ecomerce/rest/handlers/product"
	"ecomerce/rest/handlers/user"
	"ecomerce/rest/middleware"
	"fmt"
	"net/http"
	"strconv"
)

type Server struct {
	cnf            *config.Config
	productHandler *product.Handler
	userHandler    *user.Handler
}

func NewServer(cnf *config.Config,productHandler *product.Handler,
	userHandler *user.Handler,
) *Server {
	return &Server{
		cnf:cnf,
		productHandler: productHandler,
		userHandler:    userHandler,
	}
}

func (server *Server) Start() {
	manager := middleware.NewManager()
	manager.Use(
		middleware.PreFlight,
		middleware.Cors,
		middleware.Logger,
	)
	mux := http.NewServeMux()
	wrapMux := manager.WrapMux(mux)
	//InitRoute(mux, manager) // for all router
	server.productHandler.RegisterRoutes(mux, manager)
	server.userHandler.RegisterRoutes(mux, manager)
	port := ":" + strconv.Itoa(server.cnf.HttpPort)
	fmt.Println("Server is running at http://localhost:", server.cnf.HttpPort)

	err := http.ListenAndServe(port, wrapMux)
	if err != nil {
		fmt.Println("Error running server:", err)
	}
}
