package main

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"os"
)

type mockData struct {
	EventType string `json:"eventType"`
	Data      string `json:"data"`
}

func main() {
	// connect to rabbitmq
	conn, err := amqp.Dial("amqp://admin:1234@localhost:5672/")
	if err != nil {
		log.Fatalf("[main]: unable to connect RabbitMQ %+v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %s", err)
	}
	defer ch.Close()

	// publish message 10 times
	for i := 1; i <= 10; i++ {
		// data for sending each time
		body, _ := json.Marshal(&mockData{
			EventType: "sendMessage",
			Data:      fmt.Sprintf("Hello LineRookie %d", i),
		})

		// publish message
		err = ch.Publish(
			os.Getenv("EXCHANGE_NAME"),
			os.Getenv("ROUTING_KEY"),
			false,
			false,
			amqp.Publishing{
				//Priority:    uint8(i),
				ContentType: "text/plain",
				Body:        body,
			})
		if err != nil {
			log.Fatalf("Failed to publish a message: %s", err)
		}
	}

	fmt.Printf("Publish message exchange_name=%s, routing_key=%s successfully", os.Getenv("EXCHANGE_NAME"), os.Getenv("ROUTING_KEY"))
}
