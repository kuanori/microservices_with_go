package main

import (
	"context"
	"log"
	"microservices_with_go/shared/messaging"

	"github.com/rabbitmq/amqp091-go"
)

type tripConsumer struct {
	rabbitmq *messaging.RabbitMQ
}

func NewTripConsumer(rabbitmq *messaging.RabbitMQ) *tripConsumer {
	return &tripConsumer{
		rabbitmq: rabbitmq,
	}
}

func (c *tripConsumer) Listen() error {

	return c.rabbitmq.ConsumeMessages(messaging.FindAvailableDriversQueue, func(ctx context.Context, msg amqp091.Delivery) error {
		// simulate working for Fair Dispatch #10-63
		// // https://www.rabbitmq.com/tutorials/tutorial-two-go#fair-dispatch
		// time.Sleep(time.Second * 15)
		log.Printf("driver recived message: %v", msg)
		return nil
	})
}
