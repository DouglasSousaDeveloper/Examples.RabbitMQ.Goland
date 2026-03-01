package messaging

import (
	"github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	channel *amqp091.Channel
	queue   string
}

func NewConsumer(conn *amqp091.Connection, queue string) (*Consumer, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &Consumer{
		channel: ch,
		queue:   queue,
	}, nil
}

func (c *Consumer) Start(handler func([]byte)) error {
	msgs, err := c.channel.Consume(
		c.queue,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	for msg := range msgs {
		handler(msg.Body)
	}

	return nil
}