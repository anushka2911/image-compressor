package main

import (
	"fmt"

	// "github.com/anushka/consumer/pkg/imageUtils"
	"github.com/anushka/consumer/pkg/messaging"
)

func main() {
	productID, err := messaging.ConnectToRabbitMQ()
	if err != nil {
		fmt.Println("Error connecting to RabbitMQ:", err)
		return
	}

	if productID == 0 {
		fmt.Println("Error: Product ID is 0")
		return
	}

}
