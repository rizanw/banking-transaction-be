package redis

import (
	"context"
	"encoding/json"
)

func (r *rcache) Get(ctx context.Context, key string, result interface{}) error {
	serialized, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(serialized), result)
}
