package handlers

import (
	"ecomerce/database"
	"ecomerce/util"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func UpdateProduct(w http.ResponseWriter,r *http.Request){
	ProductId:=r.PathValue("id")// ProductId is a string 
	pId,err:=strconv.Atoi(ProductId)// convert string into integer which may int or not that why err

	if err!=nil{
		http.Error(w,"Please Give Me a Valid ID",http.StatusBadRequest)
		return
	}
	var newProduct database.Product
    decoder:=json.NewDecoder(r.Body)
	 err=decoder.Decode(&newProduct)
	 if err != nil {
		fmt.Println(err)
		http.Error(w, "Please Give Me a Valid Json", http.StatusBadRequest)
		return
	}
	newProduct.ID=pId
	database.Update(newProduct)
	util.SendData(w,"Successfully Updated data",http.StatusCreated)

}