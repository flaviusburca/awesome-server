package server

import (
	"container/list"
	"sync"
)

// OrderedMap is a thread-safe map that maintains the insertion order of keys.
type OrderedMap struct {
	sync.RWMutex
	data  map[string]*list.Element
	order *list.List
}

// entry is a key-value pair stored in the list.
type entry struct {
	key   string
	value string
}

func NewOrderedMap() *OrderedMap {
	return &OrderedMap{
		data:  make(map[string]*list.Element),
		order: list.New(),
	}
}

// Add adds a new key-value pair to the map. If the key already exists, the value is updated.
func (m *OrderedMap) Add(key, value string) {
	m.Lock()
	defer m.Unlock()
	if elem, ok := m.data[key]; ok {
		// Key already exists, update the value
		elem.Value.(*entry).value = value
	} else {
		// Insert new entry
		e := &entry{key, value}
		el := m.order.PushBack(e)
		m.data[key] = el
	}
}

// Delete removes a key-value pair from the map.
func (m *OrderedMap) Delete(key string) {
	m.Lock()
	defer m.Unlock()
	if elem, ok := m.data[key]; ok {
		m.order.Remove(elem)
		delete(m.data, key)
	}
}

// Get returns the value associated with the given key. If the key does not exist, the second return value is false.
func (m *OrderedMap) Get(key string) (string, bool) {
	m.RLock()
	defer m.RUnlock()
	elem, ok := m.data[key]
	if !ok {
		var zero string
		return zero, false
	}
	return elem.Value.(*entry).value, true
}

// GetAll returns a copy of the map.
func (m *OrderedMap) GetAll() []entry {
	m.RLock()
	defer m.RUnlock()
	entries := make([]entry, 0, m.order.Len())
	for elem := m.order.Front(); elem != nil; elem = elem.Next() {
		e := elem.Value.(*entry)
		entries = append(entries, *e)
	}
	return entries
}

// Keys returns a slice of all keys in insertion order.
func (m *OrderedMap) Keys() []string {
	m.RLock()
	defer m.RUnlock()

	keys := make([]string, 0, m.order.Len())
	for elem := m.order.Front(); elem != nil; elem = elem.Next() {
		keys = append(keys, elem.Value.(*entry).key)
	}
	return keys
}

// Values returns a slice of all values in insertion order.
func (m *OrderedMap) Values() []string {
	m.RLock()
	defer m.RUnlock()

	values := make([]string, 0, m.order.Len())
	for elem := m.order.Front(); elem != nil; elem = elem.Next() {
		values = append(values, elem.Value.(*entry).value)
	}
	return values
}

// Len returns the number of key-value pairs in the map.
func (m *OrderedMap) Len() int {
	m.RLock()
	defer m.RUnlock()
	return m.order.Len()
}
