package server

import (
	"reflect"
	"testing"
)

func TestAddGet(t *testing.T) {
	om := NewOrderedMap()
	om.Add("one", "1")
	om.Add("two", "2")
	om.Add("three", "3")

	if val, ok := om.Get("one"); !ok || val != "1" {
		t.Errorf("Expected value 1 for key 'one', got %v", val)
	}

	if val, ok := om.Get("two"); !ok || val != "2" {
		t.Errorf("Expected value 2 for key 'two', got %v", val)
	}

	if val, ok := om.Get("three"); !ok || val != "3" {
		t.Errorf("Expected value 3 for key 'three', got %v", val)
	}

	if val, ok := om.Get("four"); ok || val != "" {
		t.Errorf("Expected empty for key 'four', got %v", val)
	}
}

func TestAddUpdate(t *testing.T) {
	om := NewOrderedMap()
	om.Add("one", "10")
	om.Add("two", "2")
	om.Add("three", "3")

	om.Add("one", "10")

	if val, ok := om.Get("one"); !ok || val != "10" {
		t.Errorf("Expected updated value 10 for key 'one', got %v", val)
	}
}

func TestDelete(t *testing.T) {
	om := NewOrderedMap()
	om.Add("one", "1")
	om.Add("two", "2")
	om.Add("three", "3")
	om.Delete("two")
	if _, ok := om.Get("two"); ok {
		t.Error("Expected key 'two' to be deleted")
	}
}

func TestKeysAndValues(t *testing.T) {
	om := NewOrderedMap()
	om.Add("one", "1")
	om.Add("two", "2")
	om.Add("three", "3")
	om.Delete("two")
	expectedKeys := []string{"one", "three"}
	if keys := om.Keys(); !reflect.DeepEqual(keys, expectedKeys) {
		t.Errorf("Expected keys %v, got %v", expectedKeys, keys)
	}
	expectedValues := []string{"1", "3"}
	if values := om.Values(); !reflect.DeepEqual(values, expectedValues) {
		t.Errorf("Expected values %v, got %v", expectedValues, values)
	}
}

func TestOrder(t *testing.T) {
	om := NewOrderedMap()
	om.Add("one", "10")
	om.Add("two", "2")
	om.Add("three", "133")
	om.Delete("two")
	om.Add("four", "42")
	om.Add("five", "5")
	expectedKeys := []string{"one", "three", "four", "five"}
	expectedValues := []string{"10", "133", "42", "5"}
	expectedEntries := []entry{
		{"one", "10"},
		{"three", "133"},
		{"four", "42"},
		{"five", "5"},
	}
	if keys := om.Keys(); !reflect.DeepEqual(keys, expectedKeys) {
		t.Errorf("Expected keys %v, got %v", expectedKeys, keys)
	}
	if values := om.Values(); !reflect.DeepEqual(values, expectedValues) {
		t.Errorf("Expected values %v, got %v", expectedValues, values)
	}
	if allEntries := om.GetAll(); !reflect.DeepEqual(allEntries, expectedEntries) {
		t.Errorf("Expected entries %v, got %v", expectedEntries, allEntries)
	}
}
