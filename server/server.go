package server

import (
	"encoding/json"
	"github.com/flaviusburca/awesomeProject/common"
	"github.com/rabbitmq/amqp091-go"
	"log"
)

type Server struct {
	queue          common.Queue
	commandHandler CommandHandler
}

func NewServer(queue common.Queue, commandHandler CommandHandler) *Server {
	return &Server{
		queue:          queue,
		commandHandler: commandHandler,
	}
}

// Start starts to handle messages from the queue
func (s *Server) Start() {
	msgs, err := s.queue.Consume()
	if err != nil {
		log.Fatalf("Failed to retrieve messages from queue: %s", err)
	}

	for msg := range msgs {
		go s.handleMessage(msg)
	}
}

// HandleMessage handles a message from the queue
func (s *Server) handleMessage(msg amqp091.Delivery) {
	var command common.Command
	if err := json.Unmarshal(msg.Body, &command); err != nil {
		log.Printf("Error parsing message: %s", err)
		return
	}

	switch command.Type {
	case common.AddItem:
		s.commandHandler.AddItem(command.Key, command.Value)
		break
	case common.DeleteItem:
		s.commandHandler.DeleteItem(command.Key)
		break
	case common.GetItem:
		s.commandHandler.GetItem(command.Key)
		break
	case common.GetAllItems:
		s.commandHandler.GetAllItems()
		break
	default:
		log.Printf("Unknown command type: %s", command.Type)
	}
}
