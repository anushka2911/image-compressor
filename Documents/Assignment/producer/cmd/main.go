package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/streadway/amqp"

	"github.com/anushka/producer/pkg/routes"
)

type AddTask struct {
	Number1 int
	Number2 int
}

func main() {
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Printf("Failed to connect to RabbitMQ: %v", err)
		panic(err)
	}
	defer connection.Close()

	channel, err := connection.Channel()
	if err != nil {
		fmt.Printf("Failed to open a channel: %v", err)
		panic(err)
	}
	defer channel.Close()

	queue, err := channel.QueueDeclare("add", true, false, false, false, nil)
	if err != nil {
		fmt.Printf("Failed to declare a queue: %v", err)
		panic(err)
	}

	rand.Seed(time.Now().UnixNano())
	addTask := AddTask{
		Number1: rand.Intn(100),
		Number2: rand.Intn(100),
	}

	body, err := json.Marshal(addTask)
	if err != nil {
		fmt.Printf("Failed to marshal JSON in producer: %v", err)
	}

	err = channel.Publish("", queue.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         body,
	})
	if err != nil {
		fmt.Printf("Failed to publish a message: %v", err)
		panic(err)
	}

	r := mux.NewRouter()
	routes.RegisterRoutes(r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
