package kvstoreservice

import (
	"context"
	"fmt"
)

func (s *kvStoreService) Delete(ctx context.Context, key string) error {
	select {
	case <-ctx.Done():
		return fmt.Errorf("kvstoreservice.Delete storage.Delete err: %w", ctx.Err())
	default:
		if err := s.storage.Delete(key); err != nil {
			return fmt.Errorf("kvstoreservice.Delete storage.Delete err: %w", err)
		}
		return nil
	}
}
