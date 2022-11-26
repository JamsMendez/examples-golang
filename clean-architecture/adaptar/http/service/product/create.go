package product

import (
	"clean-architecture/core/dto"
	"encoding/json"
	"net/http"
)

func (service service) Create(response http.ResponseWriter, request *http.Request) {
	productRequest, err := dto.ParseJSONToCreateProductRequest(request.Body)
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	product, err := service.usecase.Create(productRequest)
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(response).Encode(product)
}
