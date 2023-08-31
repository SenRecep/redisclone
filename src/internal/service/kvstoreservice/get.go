package kvstoreservice

import (
	"context"
	"fmt"
)

func (s *kvStoreService) Get(ctx context.Context, key string) (*ItemResponse, error) {
	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("kvstoreservice.Get storage.Get err: %w", ctx.Err())
	default:
		value, err := s.storage.Get(key)
		if err != nil {
			return nil, fmt.Errorf("kvstoreservice.Get storage.Get err: %w", err)
		}
		return &ItemResponse{
			Key:   key,
			Value: value,
		}, nil
	}
}
