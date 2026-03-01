package messaging

import (
	"context"
	"encoding/json"

	"github.com/rabbitmq/amqp091-go"
)

type Producer struct {
	channel  *amqp091.Channel
	exchange string
}

func NewProducer(conn *amqp091.Connection, exchange string) (*Producer, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &Producer{
		channel:  ch,
		exchange: exchange,
	}, nil
}

func (p *Producer) Publish(routingKey string, message any) error {
	body, err := json.Marshal(message)
	if err != nil {
		return err
	}

	return p.channel.PublishWithContext(
		context.Background(),
		p.exchange,
		routingKey,
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}