package dto

import (
	"net/http"
	"strconv"
	"strings"
)

// PagiantionRequestParams is an representation query string params to filter and paginate products
type PaginationRequestParams struct {
	Search       string   `json:"search"`
	Descending   []string `json:"descending"`
	Page         int      `json:"page"`
	ItemsPerPage int      `json:"itemsPerPage"`
	Sort         []string `json:"sort"`
}

// ParseValuePaginationRequestParams converts query string params to ParseValuePaginationRequestParams struct
func ParseValuePaginationRequestParams(request *http.Request) (*PaginationRequestParams, error) {
	page, _ := strconv.Atoi(request.FormValue("page"))
	itemsPerPage, _ := strconv.Atoi(request.FormValue("itemsPerPage"))

	paginationRequestParams := PaginationRequestParams{
		Search:       request.FormValue("search"),
		Descending:   strings.Split(request.FormValue("descending"), ","),
		Sort:         strings.Split(request.FormValue("sort"), ","),
		Page:         page,
		ItemsPerPage: itemsPerPage,
	}

	return &paginationRequestParams, nil
}
