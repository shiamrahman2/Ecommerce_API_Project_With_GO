package handlers

import (
	"ecomerce/database"
	"ecomerce/util"
	"net/http"
	"strconv"
)

func GetProduct(w http.ResponseWriter,r* http.Request){
	ProductId:=r.PathValue("id")// ProductId is a string 
	pId,err:=strconv.Atoi(ProductId)// convert string into integer which may int or not that why err

	if err!=nil{
		http.Error(w,"Please Give Me a Valid ID",http.StatusBadRequest)
		return
	}

	product:=database.Get(pId)

	if product==nil{
		util.SendError(w,http.StatusNotFound,"Product doesn't found")
		return;
	}
    util.SendData(w,product,http.StatusOK)

}