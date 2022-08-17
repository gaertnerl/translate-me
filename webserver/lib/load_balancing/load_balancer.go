package loadbalancing

type LoadBalancer[T any] interface {
	Next() T
	Add(ressource T)
}
