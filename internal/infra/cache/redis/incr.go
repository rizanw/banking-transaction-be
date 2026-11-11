package redis

import (
	"context"
	"time"
)

func (r *rcache) Incr(ctx context.Context, key string, ttl time.Duration) (int64, error) {
	n, err := r.client.Incr(ctx, key).Result()
	if err != nil {
		return 0, err
	}
	if n == 1 {
		if err = r.client.Expire(ctx, key, ttl).Err(); err != nil {
			return 0, err
		}
	}
	return n, nil
}

func (r *rcache) Delete(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}
