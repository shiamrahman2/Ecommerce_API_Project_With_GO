package util

import "net/http"

type PaginatedData struct {
	Data       any        `json:"data"`
	Pagination Pagination `json:"pagination"`
}

type Pagination struct {
	Items      int64 `json:"items"`
	Page       int64 `json:"page"`
	Limit      int64 `json:"limit"`
	TotalPage  int64 `json:"totalPage"`
	TotalItems int64 `json:"totalItems"`
}

func SendPage(w http.ResponseWriter,data interface{},items,page,limit,totalPages,totalItems int64){
	response := PaginatedData{
		Data: data,
		Pagination: Pagination{
			Items:      items,
			Page:       page,
			Limit:      limit,
			TotalPage:  totalPages,
			TotalItems: totalItems,
		},
	}
	SendData(w, response, http.StatusOK)
}