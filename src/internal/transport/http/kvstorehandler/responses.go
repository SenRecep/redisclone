package kvstorehandler

import (
	"reflect"
	"sort"
)

// ItemResponse represents a key/value item.
type ItemResponse struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

// ListResponse represents a collection of ItemResponse.
type ListResponse []ItemResponse

// Order sorts the collection.
// isAsc: specifies whether the sorting is in ascending (true) or descending (false) order.
// orderFieldIsKey: specifies whether the sorting is based on the key (true) or value (false).
func (list ListResponse) Order(isAsc, orderFieldIsKey bool) {
	sort.SliceStable(list, func(i, j int) bool {
		var result bool

		if orderFieldIsKey {
			result = list[i].Key < list[j].Key
		} else {
			result = compareValues(list[i].Value, list[j].Value)
		}

		if !isAsc {
			result = !result
		}
		return result
	})
}

func compareValues(l, r any) bool {
	if !isSameType(l, r) {
		return false
	}
	switch left := l.(type) {
	case int:
		if right, ok := r.(int); ok {
			return left < right
		}
	case float64:
		if right, ok := r.(float64); ok {
			return left < right
		}
	case string:
		if right, ok := r.(string); ok {
			return left < right
		}
	case bool:
		if right, ok := r.(bool); ok {
			return left && !right
		}
	}

	return false
}

func isSameType(x, y any) bool {
	return reflect.TypeOf(x) == reflect.TypeOf(y)
}
