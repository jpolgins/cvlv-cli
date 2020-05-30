package cache

import (
	"github.com/go-redis/redis/v7"
	"time"
)

type redisCache struct {
	Client *redis.Client
}

func NewRedisCache(client *redis.Client) Cache {
	return &redisCache{
		Client: client,
	}
}

func (r redisCache) Set(key string, val interface{}, ttl time.Duration) {
	status := r.Client.Set(key, val, ttl)
	if status.Err() != nil {
		panic(status.Err())
	}
}

func (r redisCache) Get(k string) (interface{}, bool) {
	res, err := r.Client.Get(k).Result()
	if err != nil {
		return nil, false
	}

	return res, true
}
