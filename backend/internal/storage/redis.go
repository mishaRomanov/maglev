package storage

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

func NewRedisInstance(
	port string,
	password string,
	db int,
) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("127.0.0.1:%s", port),
		Password: password,
		DB:       db,
	})
}
