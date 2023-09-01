package kvstorage_test

import (
	"reflect"
	"testing"

	"github.com/SenRecep/redisclone/src/internal/storage/memory/kvstorage"
)

func TestList(t *testing.T) {
	key := "key"
	memoryStorage := kvstorage.MemoryDB{
		key: "value",
	}
	storage := kvstorage.New(
		kvstorage.WithMemoryDB(memoryStorage),
	)

	value := storage.List()

	if !reflect.DeepEqual(value, memoryStorage) {
		t.Error("value not equal")
	}
}
