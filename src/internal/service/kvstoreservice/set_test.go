package kvstoreservice_test

import (
	"context"
	"errors"
	"testing"

	"github.com/SenRecep/redisclone/src/internal/service/kvstoreservice"
)

func TestSetWithCancel(t *testing.T) {
	mockStorage := &mockStorage{
		setErr: errStorageSet,
	}
	kvsStoreService := kvstoreservice.New(kvstoreservice.WithStorage(mockStorage))

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	if _, err := kvsStoreService.Set(ctx, nil); !errors.Is(err, ctx.Err()) {
		t.Error("error not occurred")
	}
}

func TestSet(t *testing.T) {
	mockStorage := &mockStorage{
		setErr:   errStorageSet,
		memoryDB: map[string]any{},
	}
	kvsStoreService := kvstoreservice.New(kvstoreservice.WithStorage(mockStorage))

	setRequest := kvstoreservice.SetRequest{
		Key:   "key",
		Value: "value",
	}
	_, err := kvsStoreService.Set(context.Background(), &setRequest)
	if err == nil {
		t.Errorf("error occurred error: %v", err)
	}

	//todo Coverage problem

	setRequest = kvstoreservice.SetRequest{
		Key:   "key",
		Value: "value",
	}

	if _, err := kvsStoreService.Set(context.Background(), &setRequest); err == nil {
		t.Errorf("error occurred error: %v", err)
	}

}
