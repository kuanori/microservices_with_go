package events

import (
	"context"
	"microservices_with_go/shared/contracts"
	"microservices_with_go/shared/messaging"
)

type TripEventPublisher struct {
	rabbitmq *messaging.RabbitMQ
}

func NewTripEventPublisher(rabbitmq *messaging.RabbitMQ) *TripEventPublisher {
	return &TripEventPublisher{
		rabbitmq: rabbitmq,
	}
}

func (p *TripEventPublisher) PublishTripCreated(ctx context.Context) error {

	return p.rabbitmq.PublishMessage(ctx, contracts.TripEventCreated, "Trip has been created")

}
