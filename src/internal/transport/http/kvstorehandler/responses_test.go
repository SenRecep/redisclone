package kvstorehandler

import (
	"reflect"
	"testing"
)

func TestListResponseOrder(t *testing.T) {
	// Test ascending order by key
	list := ListResponse{
		{"key3", 3},
		{"key1", 1},
		{"key2", 2},
	}
	list.Order(true, true)

	expected := ListResponse{
		{"key1", 1},
		{"key2", 2},
		{"key3", 3},
	}

	if !reflect.DeepEqual(list, expected) {
		t.Errorf("Expected ascending order by key, got: %v", list)
	}

	// Test descending order by value
	list = ListResponse{
		{"key1", 10},
		{"key2", 5},
		{"key3", 20},
	}
	list.Order(false, false)

	expected = ListResponse{
		{"key3", 20},
		{"key1", 10},
		{"key2", 5},
	}

	if !reflect.DeepEqual(list, expected) {
		t.Errorf("Expected descending order by value, got: %v", list)
	}
}

func TestCompareValues(t *testing.T) {
	// Test integer comparison
	result := compareValues(5, 10)
	if !result {
		t.Errorf("Expected true, got false for integer comparison")
	}

	// Test float comparison
	result = compareValues(5.5, 10.5)
	if !result {
		t.Errorf("Expected true, got false for integer comparison")
	}

	// Test bool comparison
	result = compareValues(true, false)
	if !result {
		t.Errorf("Expected true, got false for integer comparison")
	}

	// Test string comparison
	result = compareValues("apple", "banana")
	if !result {
		t.Errorf("Expected true, got false for string comparison")
	}

	// Test struct comparison
	result = compareValues(struct{}{}, struct{}{})
	if result {
		t.Errorf("Expected true, got false for string comparison")
	}

	// Test different types comparison
	result = compareValues("apple", 10)
	if result {
		t.Errorf("Expected true, got false for string comparison")
	}
}

func TestIsSameType(t *testing.T) {
	// Test same type
	result := isSameType(10, 5)
	if !result {
		t.Errorf("Expected true, got false for same type")
	}

	// Test different types
	result = isSameType(10, "hello")
	if result {
		t.Errorf("Expected false, got true for different types")
	}
}
