package go_dingtalk_sdk_wrapper

import (
	"github.com/redis/go-redis/v9"
)

type Cache interface {
	Set()
	Get() (string, error)
}

type RedisCache struct {
	redisClient *redis.Client
}

func NewRedisCache(client *redis.Client) *RedisCache {
	return &RedisCache{
		redisClient: client,
	}
}

func (r *RedisCache) Set() error {
	return nil
}

func (r *RedisCache) Get() error {
	return nil
}
