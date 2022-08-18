package loadbalancing

import (
	"math/rand"
	"sync"
)

type RandomLB[T any] struct {
	mtx  sync.RWMutex
	list []T
}

func New_RandomLB[T any](list ...T) *RandomLB[T] {
	rlb := new(RandomLB[T])
	rlb.list = list
	return rlb
}

func (rlb *RandomLB[T]) Next() T {
	defer rlb.mtx.RUnlock()
	rlb.mtx.RLock()
	idx := rand.Intn(len(rlb.list))
	return rlb.list[idx]
}

func (rlb *RandomLB[T]) Add(ressource T) {
	defer rlb.mtx.Unlock()
	rlb.mtx.Lock()
	rlb.list = append(rlb.list, ressource)
}
