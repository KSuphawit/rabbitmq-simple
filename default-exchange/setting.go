package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

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

	args := make(amqp.Table)
	args["x-max-priority"] = 10

	// declare queue
	_, err = ch.QueueDeclare(
		"line_rookie",
		false,
		false,
		false,
		false,
		args,
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %s", err)
	}

	fmt.Println("Default exchange queue=line_rookie created")
}
