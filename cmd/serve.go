package cmd

import (
	"ecomerce/config"
	"ecomerce/infra/db"
	"ecomerce/repo"
	"ecomerce/rest"
	"ecomerce/rest/handlers/product"
	"ecomerce/rest/handlers/user"
	"ecomerce/rest/middleware"
	"fmt"
	"os"
)

func Serve() {
	cnf := config.GetConfig()
	dbCon, err := db.NewConnection()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	productRepo:=repo.NewProductRepo()
	userRepo:=repo.NewUserRepo(dbCon)
	middleware := middleware.NewMiddleWare(cnf)
	productHandler := product.NewHandler(middleware, productRepo)
	userHandler := user.NewHandler(cnf, userRepo)
	serve := rest.NewServer(cnf, productHandler, userHandler)
	serve.Start()
}
