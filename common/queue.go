package common

import "github.com/rabbitmq/amqp091-go"

// Queue is an interface for an AMQP queue
type Queue interface {
	SendCommand(command Command) error
	Consume() (<-chan amqp091.Delivery, error)
	Close()
}
