package product

import (
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
	err = h.svc.Delete(pId)
	if err != nil {
		util.SendError(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	util.SendData(w, "Successfully Deleted Data", http.StatusCreated)

}
