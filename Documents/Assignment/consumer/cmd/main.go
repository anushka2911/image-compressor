package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/streadway/amqp"
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

	messageChannel, err := channel.Consume(queue.Name, "", false, false, false, false, nil)
	if err != nil {
		fmt.Printf("Failed to register a consumer: %v", err)
		panic(err)
	}

	stopChan := make(chan bool)
	go func() {
		log.Printf("Consumer is ready,The process id PID: %d", os.Getpid())
		for d := range messageChannel {
			log.Printf("Received a message: %s", d.Body)
			addTask := AddTask{}
			err := json.Unmarshal(d.Body, addTask)
			if err != nil {
				log.Printf("Error decoding JSON while checking messages of a queue: %s", err)
			}
			log.Printf("Received a message body: %s", d.Body)
			log.Printf("Received a message: %d", addTask.Number1)
			log.Printf("Received a message: %d", addTask.Number2)
			log.Printf("Result of %d + %d is %d", addTask.Number1, addTask.Number2, addTask.Number1+addTask.Number2)

			err = d.Ack(false)
			if err != nil {
				log.Printf("Error acknowledging message : %s", err)
			} else {
				log.Printf("Acknowledged message")
			}
		}
	}()
	<-stopChan

}
