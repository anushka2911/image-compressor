package models

import (
	"fmt"
	"time"

	"github.com/anushka/producer/pkg/config"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Mobile    string    `json:"mobile"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Product struct {
	ID                      int       `json:"id"`
	UserID                  int       `json:"user_id"`
	ProductName             string    `json:"product_name"`
	ProductDescription      string    `json:"product_description"`
	ProductImages           string    `json:"product_images"`
	ProductPrice            float64   `json:"product_price"`
	CompressedProductImages string    `json:"compressed_product_images"`
	CreatedAt               time.Time `json:"created_at"`
	UpdatedAt               time.Time `json:"updated_at"`
}

func init() {
	config.Connect()
	db := config.GetDB()
	db.AutoMigrate(&User{}, &Product{})
}

func CreateProduct(userID int, productName, productDescription string, productImages string, productPrice float64) (*Product, error) {
	db := config.GetDB()
	currentTime := time.Now()

	product := &Product{
		UserID:             userID,
		ProductName:        productName,
		ProductDescription: productDescription,
		ProductImages:      productImages,
		ProductPrice:       productPrice,
		CreatedAt:          currentTime,
		UpdatedAt:          currentTime,
	}

	if err := db.Create(product).Error; err != nil {
		return nil, err
	}

	return product, nil
}

func GetProductId() (int, error) {
	db := config.GetDB()
	var product Product
	if err := db.Last(&product).Error; err != nil {
		fmt.Println("Error in fetching product id")
		return 0, err
	}

	fmt.Printf("Product ID successfully retrived: %d", product.ID)
	return product.ID, nil
}
