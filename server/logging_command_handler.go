package server

// NopCommandHandler is a decorator for StandardCommandHandler that
// overrides the GetItem and GetAllItems methods to stop creating files.
// Used for testing.
type NopCommandHandler struct {
	*StandardCommandHandler
}

func NewNopCommandHandler(orderedMap *OrderedMap) *NopCommandHandler {
	return &NopCommandHandler{
		NewStandardCommandHandler(orderedMap),
	}
}

func (h *NopCommandHandler) GetItem(key string) {
	h.orderedMap.Get(key)
}

func (h *NopCommandHandler) GetAllItems() {
	h.orderedMap.GetAll()
}
