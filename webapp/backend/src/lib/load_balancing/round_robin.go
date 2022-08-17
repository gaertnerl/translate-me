package loadbalancing

import "sync"

type RoundRobin[T any] struct {
	mtx      sync.Mutex
	Position int32
	list     []T
}

func New_RoundRobin[T any](list ...T) *RoundRobin[T] {
	rr := new(RoundRobin[T])
	rr.Position = 0
	rr.list = list
	return rr
}

func (rr *RoundRobin[T]) Next() T {
	defer rr.mtx.Unlock()
	rr.mtx.Lock()
	element := rr.list[rr.Position]
	rr.Position = (rr.Position + 1) % int32(len(rr.list))
	return element
}

func (rr *RoundRobin[T]) Add(ressource T) {
	defer rr.mtx.Unlock()
	rr.mtx.Lock()
	rr.list = append(rr.list, ressource)
}
