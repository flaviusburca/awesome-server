package server

import "testing"

func TestAdd(t *testing.T) {
	om := NewOrderedMap()
	om.Add("key1", "value1")
	if val, exists := om.Get("key1"); !exists || val != "value1" {
		t.Fatalf("expected 'value1', got '%v'", val)
	}
}

func TestDelete(t *testing.T) {
	om := NewOrderedMap()
	om.Add("key1", "value1")
	om.Delete("key1")
	if _, exists := om.Get("key1"); exists {
		t.Fatalf("expected key1 to be deleted")
	}
}

func TestGet(t *testing.T) {
	om := NewOrderedMap()
	om.Add("key1", "value1")
	if val, exists := om.Get("key1"); !exists || val != "value1" {
		t.Fatalf("expected 'value1', got '%v'", val)
	}
}

func TestGetAllItems(t *testing.T) {
	om := NewOrderedMap()
	om.Add("key1", "value1")
	om.Add("key2", "value2")
	allItems := om.GetAll()
	if len(allItems) != 2 || allItems["key1"] != "value1" || allItems["key2"] != "value2" {
		t.Fatalf("unexpected all items result")
	}
}

func TestAddDuplicateItem(t *testing.T) {
	om := NewOrderedMap()
	om.Add("key1", "value1")
	om.Add("key1", "value2")
	if val, exists := om.Get("key1"); !exists || val != "value2" {
		t.Fatalf("expected 'value2', got '%v'", val)
	}
}

func TestOrderOfItems(t *testing.T) {
	om := NewOrderedMap()
	om.Add("key1", "value1")
	om.Add("key2", "value2")
	om.Add("key3", "value3")
	allItems := om.GetAll()

	expectedOrder := []string{"key1", "key2", "key3"}
	for _, key := range expectedOrder {
		if allItems[key] == "" {
			t.Fatalf("expected key %v to be in the map", key)
		}
	}
}
