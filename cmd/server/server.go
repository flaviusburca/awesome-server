package main

import (
	"fmt"
	"github.com/flaviusburca/awesomeProject/common"
	"github.com/flaviusburca/awesomeProject/server"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := &cli.App{
		Name:   "Awesome Server",
		Usage:  "server for handling messages",
		Action: server.Launch,
	}
	app.Flags = append(app.Flags, common.RabbitMQHost, common.RabbitMQPort, common.RabbitMQUser, common.RabbitMQPassword, common.RabbitMQQueueName)

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
