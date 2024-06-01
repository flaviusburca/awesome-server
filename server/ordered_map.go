package server

import "sync"

type OrderedMap struct {
	sync.RWMutex
	data map[string]string
	keys []string
}

func NewOrderedMap() *OrderedMap {
	return &OrderedMap{
		data: make(map[string]string),
		keys: []string{},
	}
}

// Add adds a new key-value pair to the map. If the key already exists, the value is updated.
func (m *OrderedMap) Add(key, value string) {
	m.Lock()
	defer m.Unlock()
	if _, exists := m.data[key]; !exists {
		m.keys = append(m.keys, key)
	}
	m.data[key] = value
}

// Delete removes a key-value pair from the map.
func (m *OrderedMap) Delete(key string) {
	m.Lock()
	defer m.Unlock()
	if _, exists := m.data[key]; exists {
		delete(m.data, key)
		for i, k := range m.keys {
			if k == key {
				m.keys = append(m.keys[:i], m.keys[i+1:]...)
				break
			}
		}
	}
}

// Get returns the value associated with the given key. If the key does not exist, the second return value is false.
func (m *OrderedMap) Get(key string) (string, bool) {
	m.RLock()
	defer m.RUnlock()
	val, exists := m.data[key]
	return val, exists
}

// GetAll returns a copy of the map.
func (m *OrderedMap) GetAll() map[string]string {
	m.RLock()
	defer m.RUnlock()
	result := make(map[string]string)
	for _, key := range m.keys {
		result[key] = m.data[key]
	}
	return result
}
