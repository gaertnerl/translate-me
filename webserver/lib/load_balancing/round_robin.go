package loadbalancing

import "sync"

type RoundRobin[T any] struct {
	mtx      sync.Mutex
	Position int32
	List     []T
}

func New_RoundRobin[T any](list ...T) *RoundRobin[T] {
	rr := new(RoundRobin[T])
	rr.Position = 0
	rr.List = list
	return rr
}

func (rr *RoundRobin[T]) Next() T {
	rr.mtx.Lock()
	defer rr.mtx.Unlock()
	element := rr.List[rr.Position]
	rr.Position = (rr.Position + 1) % int32(len(rr.List))
	return element
}
