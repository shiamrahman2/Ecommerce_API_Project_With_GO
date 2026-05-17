package product

import (
	"ecomerce/database"
	"ecomerce/util"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	ProductId := r.PathValue("id")      // ProductId is a string
	pId, err := strconv.Atoi(ProductId) // convert string into integer which may int or not that why err

	if err != nil {
		http.Error(w, "Please Give Me a Valid ID", http.StatusBadRequest)
		return
	}
	database.Delete(pId)

	util.SendData(w, "Successfully Deleted Data", 201)

}
