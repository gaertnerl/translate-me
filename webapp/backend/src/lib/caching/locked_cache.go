package caching

import (
	"sync"
	"time"
)

type LockedCache[T Cachable] struct {
	maxSize int32
	mtx     sync.Mutex
	store   map[string]T
}

func New_LockedCache[T Cachable](maxSize int32) *LockedCache[T] {
	c := new(LockedCache[T])
	c.store = map[string]T{}
	c.maxSize = maxSize
	return c
}

func (c *LockedCache[T]) Remove(key string) {
	defer c.mtx.Unlock()
	c.mtx.Lock()
	delete(c.store, key)
}

func (c *LockedCache[T]) Add(key string, item T) {
	defer c.mtx.Unlock()
	c.mtx.Lock()
	if len(c.store) > int(c.maxSize) {
		c.store = map[string]T{}
	}
	c.store[key] = item
}

func (c *LockedCache[T]) _get(key string) (T, bool) {
	defer c.mtx.Unlock()
	c.mtx.Lock()
	val, exists := c.store[key]
	return val, exists
}

func (c *LockedCache[T]) Get(key string) (T, bool) {

	val, exists := c._get(key)

	if !exists {
		return val, exists
	}

	current := time.Now()
	if val.Timeout().UnixNano() <= current.UnixNano() {
		c.Remove(key)
		var empty T
		return empty, false
	}

	return val, exists
}
