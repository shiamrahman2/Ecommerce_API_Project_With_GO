package product

import (
	"ecomerce/domain"
	"ecomerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

type ReqProduct struct {
	Tittle      string  `json:"tittle"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgURL      string  `json:"imageURL"`
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	/*
	 1. receive body information(description,tittle,price,imageURL) from r.Body
	 2. create a instance of Product
	 3. append the instance product with productList

	*/

	var newProduct ReqProduct
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please Give Me a Valid Json", http.StatusBadRequest)
		return
	}

	createProduct, err := h.svc.Create(domain.Product{
		Tittle:      newProduct.Tittle,
		Description: newProduct.Description,
		Price:       newProduct.Price,
		ImgURL:      newProduct.ImgURL,
	})
	if err != nil {
		util.SendError(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	util.SendData(w, createProduct, http.StatusCreated)
}
