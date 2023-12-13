// db/redis.go

package db

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var REDIS *redis.Client

func ConnectToRedis(addr, password string) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return nil, fmt.Errorf("an error occurred during redis connection: %w", err)
	}

	fmt.Println("Server connected to redis")
	return rdb, nil
}
