package messaging

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/anushka/consumer/pkg/imageUtils"

	"github.com/streadway/amqp"
)

func ConnectToRabbitMQ() (int, error) {
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return 0, fmt.Errorf("failed to connect to RabbitMQ: %v", err)
	}
	defer connection.Close()

	channel, err := connection.Channel()
	if err != nil {
		return 0, fmt.Errorf("failed to open a channel: %v", err)
	}
	defer channel.Close()

	queue, err := channel.QueueDeclare("add", true, false, false, false, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to declare a queue: %v", err)
	}

	messageChannel, err := channel.Consume(queue.Name, "", false, false, false, false, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to register a consumer: %v", err)
	}

	stopChan := make(chan bool)
	var productID int

	go func() {
		log.Printf("Consumer is ready, The process id PID: %d", os.Getpid())
		for d := range messageChannel {
			log.Printf("Received a message body: %s", d.Body)

			productID, err = strconv.Atoi(string(d.Body))
			if err != nil {
				log.Printf("Error converting string to int: %s", err)
				continue
			}

			log.Printf("Received a message: %d", productID)
			//download image
			resp, err := imageUtils.DownloadAndCompressProductImages(productID)
			if err != nil {
				log.Printf("image successfully compressed and saved")
			}
			log.Printf("Response: %s", resp)
			//todo handle error here
			err = d.Ack(false)
			if err != nil {
				log.Printf("Error acknowledging message: %s", err)
			} else {
				log.Printf("Acknowledged message")
			}
		}
	}()

	<-stopChan
	return productID, nil
}
