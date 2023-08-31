package kvstorage

import (
	"fmt"
	"github.com/SenRecep/redisclone/src/internal/kverror"
)

func (ms *memoryStorage) Set(key string, value any) (any, error) {
	if _, err := ms.Get(key); err == nil { // can not set! key already exist
		return nil, fmt.Errorf("%w", kverror.ErrKeyExists.AddData("'"+key+"' exist"))
	}
	return ms.set(key, value), nil
}

func (ms *memoryStorage) set(key string, value any) any {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	ms.db[key] = value
	return value
}
