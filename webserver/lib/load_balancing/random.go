package loadbalancing

import "math/rand"

type RandomLB[T any] struct {
	list []T
}

func New_RandomLB[T any](list ...T) *RandomLB[T] {
	rlb := new(RandomLB[T])
	rlb.list = list
	return rlb
}

func (rlb *RandomLB[T]) Next() T {
	idx := rand.Intn(len(rlb.list))
	return rlb.list[idx]
}
