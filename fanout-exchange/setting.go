package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"os"
)

func main() {
	exchangeName := "youtube.notification.exchange"

	// connect to rabbitmq
	conn, err := amqp.Dial("amqp://admin:1234@localhost:5672/")
	if err != nil {
		log.Fatalf("[main]: unable to connect RabbitMQ %+v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %s", err)
	}

	// declare exchange
	err = ch.ExchangeDeclare(
		exchangeName,        // Exchange name
		amqp.ExchangeFanout, // Exchange type
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare an exchange: %s", err)
	}

	// declare queue
	q, err := ch.QueueDeclare(
		os.Getenv("QUEUE"), // Queue name (empty means a randomly generated name)
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %s", err)
	}

	// bind queue to exchange
	err = ch.QueueBind(
		q.Name, // Queue name
		"",     // Routing key (empty for fanout exchange type)
		exchangeName,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to bind a queue to the exchange: %s", err)
	}

	fmt.Printf("Fanout exchange queue=%s routing_key='' binding_exchange=%s created", os.Getenv("QUEUE"), exchangeName)
}
