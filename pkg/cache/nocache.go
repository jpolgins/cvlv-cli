package cache

import (
	"time"
)

type inMemory struct{}

func NewNoCache() Cache {
	return &inMemory{}
}

func (inMemory) Set(key string, val interface{}, duration time.Duration) {
}

func (inMemory) Get(k string) (interface{}, bool) {
	return nil, false
}

