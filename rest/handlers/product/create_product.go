package product

import (
	"ecomerce/database"
	"ecomerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

func(h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	/*
	 1. receive body information(description,tittle,price,imageURL) from r.Body
	 2. create a instance of Product
	 3. append the instance product with productList

	*/

	var newProduct database.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please Give Me a Valid Json", http.StatusBadRequest)
		return
	}

	createProduct := database.Store(newProduct)

	util.SendData(w, createProduct, 201)
}
