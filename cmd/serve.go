package cmd

import (
	"ecomerce/config"
	"ecomerce/infra/db"
	"ecomerce/product"
	"ecomerce/repo"
	"ecomerce/rest"
	productHandler "ecomerce/rest/handlers/product"
	userHandler "ecomerce/rest/handlers/user"
	"ecomerce/rest/middleware"
	"ecomerce/user"
	"fmt"
	"os"
)

func Serve() {
	cnf := config.GetConfig()
	dbCon, err := db.NewConnection(cnf.DB)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err=db.MigrateDB(dbCon,"./migrations")
	if err!=nil{
		fmt.Println(err)
		os.Exit(1)
	}
	// Repository
	productRepo:=repo.NewProductRepo(dbCon)
	userRepo:=repo.NewUserRepo(dbCon)
    // domain
	userSVC:=user.NewService(userRepo)
    prdSVC:=product.NewService(productRepo)
	middleware := middleware.NewMiddleWare(cnf)
	productHandler := productHandler.NewHandler(middleware, prdSVC)
	userHandler := userHandler.NewHandler(cnf, userSVC)
	serve := rest.NewServer(cnf, productHandler, userHandler)
	serve.Start()
}
