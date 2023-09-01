package kvstorage_test

import (
	"testing"

	"github.com/SenRecep/redisclone/src/internal/storage/memory/kvstorage"
)

func TestSet(t *testing.T) {
	key := "key"
	memoryStorage := kvstorage.MemoryDB{}
	storage := kvstorage.New(
		kvstorage.WithMemoryDB(memoryStorage),
	)

	_, err := storage.Set(key, "value")
	if err != nil {
		t.Error(err)
	}

	if _, ok := memoryStorage[key]; !ok {
		t.Error("value not equal")
	}
}
