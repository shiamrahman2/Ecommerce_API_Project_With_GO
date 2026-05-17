package product

import (
	"ecomerce/database"
	"ecomerce/util"
	"net/http"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	util.SendData(w,database.List(),http.StatusOK)
}
