package server

import (
	"fmt"
	"log"
	"os"
	"time"
)

// CommandHandler defines the supported commands
type CommandHandler interface {
	AddItem(key string, value string)
	DeleteItem(key string)
	GetItem(key string)
	GetAllItems()
}

// StandardCommandHandler is the default implementation of CommandHandler
type StandardCommandHandler struct {
	orderedMap *OrderedMap
}

// NewStandardCommandHandler creates a new StandardCommandHandler
func NewStandardCommandHandler(orderedMap *OrderedMap) *StandardCommandHandler {
	return &StandardCommandHandler{
		orderedMap: orderedMap,
	}
}

// AddItem adds a new key-value pair to the map. If the key already exists, the value is updated.
func (h *StandardCommandHandler) AddItem(key string, value string) {
	h.orderedMap.Add(key, value)
}

// DeleteItem removes a key-value pair from the map.
func (h *StandardCommandHandler) DeleteItem(key string) {
	h.orderedMap.Delete(key)
}

// GetItem writes the value associated with the given key to a file. If the key does not exist, nothing is done.
func (h *StandardCommandHandler) GetItem(key string) {
	value, exists := h.orderedMap.Get(key)
	if exists {
		err := os.WriteFile(
			fmt.Sprintf("output-%d.txt", time.Now().UnixMilli()),
			[]byte(fmt.Sprintf("%s:%s", key, value)),
			0644,
		)
		if err != nil {
			log.Printf("Error writing file: %s", err)
		}
	}
}

// GetAllItems writes all the key-value pairs to a file.
func (h *StandardCommandHandler) GetAllItems() {
	allItems := h.orderedMap.GetAll()
	f, err := os.Create(fmt.Sprintf("output-%d.txt", time.Now().UnixMilli()))
	if err != nil {
		log.Printf("Error creating file: %s", err)
		_ = f.Close()
		return
	}
	for key, value := range allItems {
		_, err := fmt.Fprintf(f, "%s:%s\n", key, value)
		if err != nil {
			log.Printf("Error writing to file: %s", err)
		}
	}
	err = f.Close()
	if err != nil {
		log.Printf("Error closing file: %s", err)
	}
}
