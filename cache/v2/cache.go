package v2

import "time"

// Cache defines the generics cache interface for 1.8 or above version.
// Notice: SetWithExp function is not necessarily implemented
type Cache[K comparable, V any] interface {
	Get(key K) (value V, ok bool)
	Set(key K, val V)
	SetWithExp(key K, val V, exp time.Duration)
	Delete(key K)
}
