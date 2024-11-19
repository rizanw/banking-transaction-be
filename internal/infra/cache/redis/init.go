package redis

import (
	"tx-bank/internal/infra/cache"

	"github.com/redis/go-redis/v9"
)

type rcache struct {
	client *redis.Client
}

func New(cli *redis.Client) cache.Cache {
	return &rcache{client: cli}
}
