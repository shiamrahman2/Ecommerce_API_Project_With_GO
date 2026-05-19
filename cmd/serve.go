package cmd

import (
	repo "ecomerce/Repo"
	"ecomerce/config"
	"ecomerce/rest"
	"ecomerce/rest/handlers/product"
	"ecomerce/rest/handlers/user"
	"ecomerce/rest/middleware"
)

func Serve() {
	cnf := config.GetConfig()
	productRepo := repo.NewProductRepo()
	userRepo := repo.NewUserRepo()
	middleware := middleware.NewMiddleWare(cnf)
	productHandler := product.NewHandler(middleware, productRepo)
	userHandler := user.NewHandler(cnf, userRepo)
	serve := rest.NewServer(cnf, productHandler, userHandler)
	serve.Start()
}
