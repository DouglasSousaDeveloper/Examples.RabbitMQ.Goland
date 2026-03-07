package main

import (
	"fmt"
	"log"

	"github.com/DouglasSousaDeveloper/Examples.RabbitMQ.Goland/internal/domain"
	"github.com/DouglasSousaDeveloper/Examples.RabbitMQ.Goland/internal/messaging"
)

func main() {
	conn, err := messaging.NewConnection("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	publisher, err := messaging.NewPublisher(conn, "billing.exchange")
	if err != nil {
		log.Fatal(err)
	}

	customer := domain.GenerateFake(nil)

	err = publisher.Publish("billing.created", customer)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Customer published successfully!")
}