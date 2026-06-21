package product

import (
	"ecomerce/util"
	"net/http"
	"strconv"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	reqQuery := r.URL.Query()

	pageAsStr := reqQuery.Get("page")
	limitAsStr := reqQuery.Get("limit")

	page, _ := strconv.ParseInt(pageAsStr, 10, 64)
	limit, _ := strconv.ParseInt(limitAsStr, 10, 64)

	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = 10
	}

	productList, err := h.svc.List(page, limit)
	if err != nil {
		util.SendError(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	totalItems, err := h.svc.Count()
	if err != nil {
		util.SendError(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	totalPages := (totalItems + limit - 1) / limit
	items := int64(len(productList))

	util.SendPage(w, productList, items, page, limit, totalPages, totalItems)
}
