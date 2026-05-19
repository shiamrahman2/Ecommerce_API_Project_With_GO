package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func GetConnectionString() string {
	/*
	  user->postgres
	  password->#shiam@019#
	  host->localhost
	  port->5432

	*/

	return "user=postgres password=#shiam@019# host=localhost port=5432 dbname=ecommerce"
}

func NewConnection() (*sqlx.DB, error) {
	dbSource := GetConnectionString()
	dbCon,err:=sqlx.Connect("postgres", dbSource)
	if err!=nil{
		fmt.Println(err)
		return nil,err
	}
	return dbCon,nil
}
