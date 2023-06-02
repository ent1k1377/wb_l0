package rediscache

import (
	"github.com/go-redis/redis"
	"time"
)

type Redis struct {
	client *redis.Client
}

func New(client *redis.Client) *Redis {
	return &Redis{
		client: client,
	}
}

func (r *Redis) Set(key string, value string, duration time.Duration) error {
	err := r.client.Set(key, value, duration).Err()
	return err
}

func (r *Redis) Get(key string) (string, error) {
	value, err := r.client.Get(key).Result()
	if err != nil {
		return "", err
	}
	return value, nil
}
