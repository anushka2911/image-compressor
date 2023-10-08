package main

import (
	"fmt"

	"github.com/anushka/consumer/pkg/messaging"
)

func main() {
	err := messaging.ConnectToRabbitMQ()
	if err != nil {
		fmt.Println("Error connecting to RabbitMQ:", err)
		return
	}

}
