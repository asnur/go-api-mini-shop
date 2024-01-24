package config

import (
	"context"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

func ConnectRedis() (*redis.Client, error) {
	// Address
	address := fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))

	rdb := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	// PING
	_, err := rdb.Ping(context.TODO()).Result()

	return rdb, err
}
