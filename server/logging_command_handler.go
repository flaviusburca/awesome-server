package server

import (
	"fmt"
	"log"
	"strings"
)

// LoggingCommandHandler is a decorator for StandardCommandHandler that logs all the commands that are executed
// and overrides the GetItem and GetAllItems methods to write to the standard output.
type LoggingCommandHandler struct {
	*StandardCommandHandler
}

func NewLoggingCommandHandler(orderedMap *OrderedMap) *LoggingCommandHandler {
	return &LoggingCommandHandler{
		NewStandardCommandHandler(orderedMap),
	}
}

func (h *LoggingCommandHandler) AddItem(key string, value string) {
	log.Printf("AddItem(%s, %s)\n", key, value)
	h.StandardCommandHandler.AddItem(key, value)
}

func (h *LoggingCommandHandler) DeleteItem(key string) {
	log.Printf("DeleteItem(%s)\n", key)
	h.StandardCommandHandler.DeleteItem(key)
}

// GetItem writes the value associated with the given key to the standard output. If the key does not exist, nothing is done.
func (h *LoggingCommandHandler) GetItem(key string) {
	value, _ := h.orderedMap.Get(key)
	log.Printf("GetItem(%s): %s", key, value)
}

// GetAllItems writes all the key-value pairs to the standard output.
func (h *LoggingCommandHandler) GetAllItems() {
	allItems := h.orderedMap.GetAll()
	var sb strings.Builder
	for key, value := range allItems {
		sb.WriteString(fmt.Sprintf("%s:%s ", key, value))
	}
	if sb.Len() > 0 {
		log.Println(sb.String())
	}
}
