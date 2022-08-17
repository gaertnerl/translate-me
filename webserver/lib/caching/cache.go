package caching

import "time"

type Cachable interface {
	Timeout() time.Time
}

type Cache[T Cachable] interface {
	Get(key string) (T, bool)
	Add(key string, item T)
	Remove(Key string)
}
