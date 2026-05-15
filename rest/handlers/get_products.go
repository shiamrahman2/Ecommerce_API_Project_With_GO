package handlers

import (
	"ecomerce/database"
	"ecomerce/util"
	"net/http"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	util.SendData(w,database.List(),http.StatusOK)
}
