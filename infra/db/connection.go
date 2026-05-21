package db

import (
	"ecomerce/config"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnectionString(cnf *config.DBConfig) string {
	connectionString := fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s",
		cnf.User,
		cnf.Password,
		cnf.Host,
		cnf.Port,
		cnf.Name,
	)
	if !cnf.EnableSSLMode {
		connectionString += " sslmode=disable"
	}
	return connectionString

}

func NewConnection(cnf *config.DBConfig) (*sqlx.DB, error) {

	dbSource := GetConnectionString(cnf)

	dbCon, err := sqlx.Connect("postgres", dbSource)
	if err != nil {
		return nil, err
	}

	return dbCon, nil
}
