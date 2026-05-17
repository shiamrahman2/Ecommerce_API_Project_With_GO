package cmd

import (
	"ecomerce/config"
	"ecomerce/rest"
	"ecomerce/rest/handlers/product"
	"ecomerce/rest/handlers/user"
	"ecomerce/rest/middleware"
)

func Serve() {
	cnf := config.GetConfig()
	middleware:=middleware.NewMiddleWare(cnf)
	productHandler := product.NewHandler(middleware)
	userHandler := user.NewHandler()
	serve := rest.NewServer(cnf, productHandler, userHandler)
	serve.Start()
}
