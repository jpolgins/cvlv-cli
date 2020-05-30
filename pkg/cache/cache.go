package cache

import "time"

type Cache interface {
	Set(key string, val interface{}, duration time.Duration)
	Get(k string) (interface{}, bool)
}
