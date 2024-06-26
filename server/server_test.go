package server

import (
	"fmt"
	"github.com/flaviusburca/awesomeProject/common"
	"log"
	"testing"
)

func BenchmarkServer(b *testing.B) {
	N := 500_000
	queue, err := common.NewInMemoryQueue(2 * N)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}

	srv := NewServer(queue, NewNopCommandHandler(NewOrderedMap()))

	commands := make([]common.Command, N)
	for i := 0; i < N; i++ {
		var command common.Command
		if i%2 == 0 {
			// AddItem command
			command = common.Command{
				Type:  common.AddItem,
				Key:   "key" + fmt.Sprint(i/2),
				Value: "value" + fmt.Sprint(i/2),
			}
		} else {
			// DeleteItem command
			command = common.Command{
				Type: common.DeleteItem,
				Key:  "key" + fmt.Sprint(i/2),
			}
		}
		if i%6 == 0 {
			// GetItem command
			command = common.Command{
				Type: common.GetItem,
				Key:  "key" + fmt.Sprint(i/6),
			}
		}
		commands[i] = command
	}

	go srv.Start()

	b.ResetTimer()
	for i := 0; i < len(commands); i++ {
		err := queue.SendCommand(commands[i])
		if err != nil {
			b.Fatalf("Failed to send command: %s", err)
		}
	}
	b.StopTimer()
}
