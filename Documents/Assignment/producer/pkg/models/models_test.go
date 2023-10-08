package models

import (
	"testing"
)

func TestCreateProduct(t *testing.T) {
	product, err := CreateProduct(1, "test product", "test description", "test image", 10.0)
	if err != nil {
		t.Errorf("Error creating product: %s", err.Error())
	}

	if product.UserID != 1 {
		t.Errorf("Expected UserID to be 1, got %d", product.UserID)
	}
	if product.ProductName != "test product" {
		t.Errorf("Expected ProductName to be 'test product', got '%s'", product.ProductName)
	}
	if product.ProductDescription != "test description" {
		t.Errorf("Expected ProductDescription to be 'test description', got '%s'", product.ProductDescription)
	}
	if product.ProductImages != "test image" {
		t.Errorf("Expected ProductImages to be 'test image', got '%s'", product.ProductImages)
	}
	if product.ProductPrice != 10.0 {
		t.Errorf("Expected ProductPrice to be 10.0, got %f", product.ProductPrice)
	}
}

func TestGetProductId(t *testing.T) {
	productID, err := GetProductId()
	if err != nil {
		t.Errorf("Error getting product ID: %s", err.Error())
	}

	if productID <= 0 {
		t.Errorf("Expected product ID to be greater than 0, got %d", productID)
	}
}

func TestGetProductImagesByProductID(t *testing.T) {
	images, err := GetProductImagesByProductID(7)
	if err != nil {
		t.Errorf("Error getting product images: %s", err.Error())
	}

	if len(images) == 0 {
		t.Errorf("Expected images to not be empty")
	}
}

func BenchmarkCreateProduct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := CreateProduct(1, "test product", "test description", "test image", 10.0)
		if err != nil {
			b.Errorf("Error creating product: %s", err.Error())
		}
	}
}

func BenchmarkGetProductId(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := GetProductId()
		if err != nil {
			b.Errorf("Error getting product ID: %s", err.Error())
		}
	}
}

func BenchmarkGetProductImagesByProductID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := GetProductImagesByProductID(7)
		if err != nil {
			b.Errorf("Error getting product images: %s", err.Error())
		}
	}
}
