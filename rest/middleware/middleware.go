package middleware

import "ecomerce/config"

type MiddleWare struct {
	cnf *config.Config
}

func NewMiddleWare(cnf *config.Config) *MiddleWare {
	return &MiddleWare{
		cnf: cnf,
	}
}
