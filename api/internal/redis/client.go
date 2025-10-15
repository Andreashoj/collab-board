package redisclient

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Client struct {
	Rdb *redis.Client
}

func NewClient() *Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	ctx := context.Background()
	if err := rdb.Ping(ctx).Err(); err != nil {
		panic("Failed to connect to Redis: " + err.Error())
	}

	return &Client{Rdb: rdb}
}
