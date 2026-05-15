package cmd

import (
	"ecomerce/config"
	"ecomerce/rest"
)

func Serve() {
	cnf := config.GetConfig()
	rest.Start(cnf)
}
