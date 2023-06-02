package cache

import "time"

type Cache interface {
	Set(key string, value string, duration time.Duration) error
	Get(key string) (string, error)
}
