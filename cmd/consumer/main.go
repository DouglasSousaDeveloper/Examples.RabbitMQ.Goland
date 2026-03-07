package main

import (
	"fmt"
	"log"

	"github.com/DouglasSousaDeveloper/Examples.RabbitMQ.Goland/internal/messaging"
)

func main() {
	conn, err := messaging.NewConnection("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	consumer, err := messaging.NewConsumer(conn, "billing.queue")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Waiting for messages...")

	err = consumer.Start(func(body []byte) {
		fmt.Println("Message received:", string(body))
	})

	if err != nil {
		log.Fatal(err)
	}
}