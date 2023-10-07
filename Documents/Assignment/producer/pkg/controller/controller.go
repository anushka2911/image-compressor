package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/anushka/producer/pkg/models"
)

func CreateProduct(response http.ResponseWriter, request *http.Request) {
	var createProduct models.Product
	if err := json.NewDecoder(request.Body).Decode(&createProduct); err != nil {
		http.Error(response, "Invalid request payload", http.StatusBadRequest)
		return
	}

	product, err := models.CreateProduct(createProduct.UserID, createProduct.ProductName, createProduct.ProductDescription, createProduct.ProductImages, createProduct.ProductPrice)
	if err != nil {
		http.Error(response, "Error creating product", http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(product)
	if err != nil {
		http.Error(response, "Error encoding product to JSON", http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusOK)
	response.Write(jsonResponse)
}
