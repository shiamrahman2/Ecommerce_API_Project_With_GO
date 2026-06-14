package product

import (
	"ecomerce/domain"
	"ecomerce/util"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)
type UpdateProduct struct {
	Tittle      string  `json:"tittle"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgURL      string  `json:"imageURL"`
}

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	ProductId := r.PathValue("id")      // ProductId is a string
	pId, err := strconv.Atoi(ProductId) // convert string into integer which may int or not that why err

	if err != nil {
		http.Error(w, "Please Give Me a Valid ID", http.StatusBadRequest)
		return
	}
	var updateProduct UpdateProduct
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&updateProduct)
	if err != nil {
		fmt.Println(err)
		util.SendError(w, "Please Give Me a Valid Json", http.StatusBadRequest)
		return
	}
	_,err=h.svc.Update(domain.Product{
        ID:pId,
		Tittle:updateProduct.Tittle,
		Description:updateProduct.Description,
		ImgURL: updateProduct.ImgURL,
		Price:updateProduct.Price,
	})
	if err!=nil{
		util.SendError(w,"Internal Server Error",http.StatusInternalServerError)
		return
	}
	util.SendData(w, "Successfully Updated data", http.StatusCreated)

}
