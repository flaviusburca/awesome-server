package common

import (
	"encoding/json"
	"errors"
	"github.com/rabbitmq/amqp091-go"
	"sync"
)

// InMemoryQueue is an in-memory implementation of the Queue interface used for testing and benchmarking purposes
type InMemoryQueue struct {
	messages chan amqp091.Delivery
	closed   bool
	mutex    sync.Mutex
}

func NewInMemoryQueue(bufferSize int) (Queue, error) {
	return &InMemoryQueue{
		messages: make(chan amqp091.Delivery, bufferSize),
		closed:   false,
	}, nil
}

func (q *InMemoryQueue) SendCommand(command Command) error {
	body, err := json.Marshal(command)
	if err != nil {
		return err
	}
	q.mutex.Lock()
	defer q.mutex.Unlock()
	if q.closed {
		return errors.New("queue is closed")
	}
	q.messages <- amqp091.Delivery{Body: body}
	return nil
}

func (q *InMemoryQueue) Consume() (<-chan amqp091.Delivery, error) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	if q.closed {
		return nil, errors.New("queue is closed")
	}
	return q.messages, nil
}

func (q *InMemoryQueue) Close() {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	if !q.closed {
		close(q.messages)
		q.closed = true
	}
}
