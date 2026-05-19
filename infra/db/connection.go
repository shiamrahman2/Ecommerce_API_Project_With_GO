package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnectionString() string {

	return "user=postgres password='#shiam@019#' host=localhost port=5432 dbname=ecommerce sslmode=disable"
}

func NewConnection() (*sqlx.DB, error) {

	dbSource := GetConnectionString()

	dbCon, err := sqlx.Connect("postgres", dbSource)
	if err != nil {
		return nil, err
	}

	return dbCon, nil
}
