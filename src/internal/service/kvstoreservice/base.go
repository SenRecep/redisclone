package kvstoreservice

import (
	"context"

	"github.com/SenRecep/redisclone/src/internal/storage/memory/kvstorage"
)

var _ KVStoreService = (*kvStoreService)(nil) // compile time proof

// KVStoreService defines service behaviours.
type KVStoreService interface {
	Set(context.Context, *SetRequest) (*ItemResponse, error)
	Get(context.Context, string) (*ItemResponse, error)
	Update(context.Context, *UpdateRequest) (*ItemResponse, error)
	Delete(context.Context, string) error
	List(context.Context) (*ListResponse, error)
}

type kvStoreService struct {
	storage kvstorage.Storer
}

// ServiceOption represents service option type.
type ServiceOption func(*kvStoreService)

// WithStorage sets storage option.
func WithStorage(storage kvstorage.Storer) ServiceOption {
	return func(s *kvStoreService) {
		s.storage = storage
	}
}

// New instantiates new service instance.
func New(options ...ServiceOption) KVStoreService {
	kvs := &kvStoreService{}

	for _, o := range options {
		o(kvs)
	}

	return kvs
}
