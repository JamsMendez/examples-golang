package product

import (
	"clean-architecture/core/dto"
	"encoding/json"
	"net/http"
)

func (service service) Fetch(response http.ResponseWriter, request *http.Request) {
  paginationRequestParams, err := dto.ParseValuePaginationRequestParams(request)
  if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
  }

  products, err := service.usecase.Fetch(paginationRequestParams)
  if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
  }

  json.NewEncoder(response).Encode(products)
}
