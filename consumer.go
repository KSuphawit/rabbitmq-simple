package main

import (
	"github.com/streadway/amqp"
	"log"
	"os"
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

	// Set consume size
	//err = ch.Qos(
	//	1,
	//	0,
	//	false,
	//)

	// consume message from queue
	msgs, err := ch.Consume(
		os.Getenv("QUEUE"),
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %s", err)
	}

	forever := make(chan bool)

	go func() {
		for msg := range msgs {
			log.Printf("Received message: %s", msg.Body)
			// Manual Acknowledge the message
			//time.Sleep(3 * time.Second)
			//err := msg.Ack(false)
			//if err != nil {
			//	log.Printf("Failed to acknowledge message: %v", err)
			//}
		}
	}()

	log.Printf("Waiting for messages from queue=%s To exit, press CTRL+C", os.Getenv("QUEUE"))
	<-forever
}
