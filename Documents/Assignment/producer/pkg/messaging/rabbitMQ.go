package messaging

import (
	"encoding/json"
	"fmt"

	"github.com/streadway/amqp"
)

const (
	queueName = "processImageQueue"
)

func ConnectToRabbitMQ(productID int) error {
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return fmt.Errorf("failed to connect to RabbitMQ: %v", err)
	}
	defer connection.Close()

	channel, err := connection.Channel()
	if err != nil {
		return fmt.Errorf("failed to open a channel: %v", err)
	}
	defer channel.Close()

	queue, err := channel.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		return fmt.Errorf("failed to declare a queue: %v", err)
	}

	body, err := json.Marshal(productID)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON in producer: %v", err)
	}

	err = channel.Publish("", queue.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         body,
	})
	if err != nil {
		return fmt.Errorf("failed to publish a message: %v", err)
	}

	return nil
}
