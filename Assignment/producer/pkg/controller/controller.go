package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/anushka/producer/pkg/messaging"
	"github.com/anushka/producer/pkg/models"
	"github.com/sirupsen/logrus"
)

func respondWithError(response http.ResponseWriter, errorMsg string, statusCode int) {
	http.Error(response, errorMsg, statusCode)
}

func respondWithInternalError(response http.ResponseWriter, errorMsg string) {
	logrus.Error(errorMsg)
	respondWithError(response, "Internal Server Error: "+errorMsg, http.StatusInternalServerError)
}

func CreateProduct(response http.ResponseWriter, request *http.Request) {
	var createProduct models.Product

	if err := json.NewDecoder(request.Body).Decode(&createProduct); err != nil {
		respondWithError(response, "Invalid request payload", http.StatusBadRequest)
		return
	}

	product, err := models.CreateProduct(createProduct.UserID, createProduct.ProductName, createProduct.ProductDescription, createProduct.ProductImages, createProduct.ProductPrice)
	if err != nil {
		respondWithInternalError(response, "Error creating product")
		return
	}

	jsonResponse, err := json.Marshal(product)
	if err != nil {
		respondWithInternalError(response, "Error encoding product to JSON")
		return
	}

	productID, err := models.GetProductId()
	if err != nil {
		respondWithInternalError(response, "Error getting product id")
		return
	}
	if productID == 0 {
		respondWithInternalError(response, "Error getting product id")
		return
	}

	err = messaging.ConnectToRabbitMQ(productID)
	if err != nil {
		respondWithInternalError(response, "Error connecting to RabbitMQ")
		return
	}

	response.WriteHeader(http.StatusOK)
	response.Write(jsonResponse)
}
