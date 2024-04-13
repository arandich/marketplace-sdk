package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
)

func Connect(ctx context.Context, cfg Config) (*redis.Client, error) {
	opts := &redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	}

	rdb := redis.NewClient(opts)

	err := rdb.Ping(ctx).Err()
	if err != nil {
		return nil, err
	}

	return rdb, nil
}
