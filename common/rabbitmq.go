package common

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/rabbitmq/amqp091-go"
	"github.com/urfave/cli"
	"log"
)

var (
	RabbitMQHost = cli.StringFlag{
		Name:  "host",
		Value: "localhost",
		Usage: "RabbitMQ host",
	}
	RabbitMQPort = cli.IntFlag{
		Name:  "port",
		Usage: "RabbitMQ port",
		Value: 5672,
	}
	RabbitMQUser = cli.StringFlag{
		Name:  "user",
		Value: "guest",
		Usage: "RabbitMQ user",
	}
	RabbitMQPassword = cli.StringFlag{
		Name:  "password",
		Value: "guest",
		Usage: "RabbitMQ password",
	}
	RabbitMQQueueName = cli.StringFlag{
		Name:  "queue",
		Value: "commands",
		Usage: "RabbitMQ queue for receiving commands",
	}
)

// RabbitMQQueue is an implementation of the Queue interface using RabbitMQ
type RabbitMQQueue struct {
	Conn    *amqp091.Connection
	Channel *amqp091.Channel
	Name    string
}

func NewRabbitMQQueue(ctx *cli.Context) (Queue, error) {
	url := fmt.Sprintf("amqp://%s:%s@%s:%s/",
		ctx.String(RabbitMQUser.Name),
		ctx.String(RabbitMQPassword.Name),
		ctx.String(RabbitMQHost.Name),
		ctx.String(RabbitMQPort.Name),
	)
	conn, err := amqp091.Dial(url)
	if err != nil {
		return nil, errors.New("Failed to connect to RabbitMQ [" + url + "]: " + err.Error())
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	_, err = ch.QueueDeclare(
		ctx.String(RabbitMQQueueName.Name),
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, errors.New("Failed to declare queue " + ctx.String(RabbitMQQueueName.Name) + ":" + err.Error())
	}

	return &RabbitMQQueue{
		Conn:    conn,
		Channel: ch,
		Name:    ctx.String(RabbitMQQueueName.Name),
	}, nil
}

func (q *RabbitMQQueue) Consume() (<-chan amqp091.Delivery, error) {
	msgs, err := q.Channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	return msgs, err
}

func (q *RabbitMQQueue) SendCommand(command Command) error {
	body, err := json.Marshal(command)
	if err != nil {
		return err
	}

	err = q.Channel.Publish(
		"",
		q.Name,
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	return err
}

func (q *RabbitMQQueue) Close() {
	err := q.Channel.Close()
	if err != nil {
		log.Printf("Failed to close the RabbitMQ channel: %s", err)
	}
	err = q.Conn.Close()
	if err != nil {
		log.Printf("Failed to close the RabbitMQ connection: %s", err)
	}
}
