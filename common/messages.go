package common

// Command the command exchanged between the client and the server
type Command struct {
	Type  string
	Key   string
	Value string
}

const (
	// Supported commands
	AddItem     = "addItem"
	DeleteItem  = "deleteItem"
	GetItem     = "getItem"
	GetAllItems = "getAllItems"
)
