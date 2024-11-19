package redis

import (
	"context"
	"encoding/json"
	"time"
)

func (r *rcache) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	serialized, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return r.client.Set(ctx, key, serialized, ttl).Err()
}
