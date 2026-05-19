package product

import (
	"ecomerce/util"
	"net/http"
	"strconv"
)

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	ProductId := r.PathValue("id")      // ProductId is a string
	pId, err := strconv.Atoi(ProductId) // convert string into integer which may int or not that why err

	if err != nil {
		http.Error(w, "Please Give Me a Valid ID", http.StatusBadRequest)
		return
	}

	product, err := h.productRepo.Get(pId)
	if err != nil {
		util.SendError(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if product == nil {
		util.SendError(w, "Product doesn't found", http.StatusNotFound)
		return
	}
	util.SendData(w, product, http.StatusOK)

}
