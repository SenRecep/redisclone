package kvstoreservice

import (
	"context"
	"fmt"
)

func (s *kvStoreService) List(ctx context.Context) (*ListResponse, error) {
	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("kvstoreservice.List storage.List err: %w", ctx.Err())
	default:
		items := s.storage.List()
		response := make(ListResponse, len(items))

		var i int
		for k, v := range items {
			response[i] = ItemResponse{
				Key:   k,
				Value: v,
			}
			i++
		}
		return &response, nil
	}
}
