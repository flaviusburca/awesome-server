package server

import (
	"github.com/flaviusburca/awesomeProject/common"
	"github.com/urfave/cli"
	"log"
)

// Launch is the main entry point for the server
// Here you can configure the CommandHandler and common.Queue implementations
func Launch(ctx *cli.Context) error {
	queue, err := common.NewRabbitMQQueue(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}
	defer queue.Close()
	srv := NewServer(queue, NewStandardCommandHandler(NewOrderedMap()))
	srv.Start()

	return nil
}
