package main

import (
	"bufio"
	"fmt"
	"github.com/flaviusburca/awesomeProject/common"
	"github.com/urfave/cli"
	"log"
	"os"
	"strings"
)

func main() {
	app := &cli.App{
		Name:   "Awesome Server CLI",
		Usage:  "client for sending messages",
		Action: Launch,
	}
	app.Flags = append(app.Flags, common.RabbitMQHost, common.RabbitMQPort, common.RabbitMQUser, common.RabbitMQPassword, common.RabbitMQQueueName)

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// Launch is the main entry point for the client
func Launch(ctx *cli.Context) error {
	queue, err := common.NewRabbitMQQueue(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}
	defer queue.Close()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}

		commandType := parts[0]
		command := common.Command{Type: commandType}

		switch commandType {
		case common.AddItem:
			if len(parts) != 3 {
				log.Printf("Invalid number of arguments for command %s", commandType)
				continue
			}
			command.Key = parts[1]
			command.Value = parts[2]
			break
		case common.DeleteItem, common.GetItem:
			if len(parts) != 2 {
				log.Printf("Invalid number of arguments for command %s", commandType)
				continue
			}
			command.Key = parts[1]
			break
		case common.GetAllItems:
			// no additional arguments
		}

		err = queue.SendCommand(command)
		if err != nil {
			log.Printf("Failed to send command: %s", err)
		} else {
			log.Println("Command sent")
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return nil
}
