package cache

import (
	"context"
	"time"
)

type Cache interface {
	Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error
	Get(ctx context.Context, key string, result interface{}) error
	Delete(ctx context.Context, key string) error
	// Incr increments a counter and sets its TTL when it's created.
	Incr(ctx context.Context, key string, ttl time.Duration) (int64, error)
}
