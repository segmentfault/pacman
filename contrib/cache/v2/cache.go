package v2

import (
	"time"

	genericsCache "github.com/Code-Hex/go-generics-cache"
	v2 "github.com/segmentfault/pacman/cache/v2"
)

// Cache is implemented by third-party packages which provide from github.com/Code-Hex/go-generics-cache
type Cache[K comparable, V any] struct {
	cache *genericsCache.Cache[K, V]
}

func (c Cache[K, V]) Get(key K) (value V, ok bool) {
	return c.cache.Get(key)
}

func (c Cache[K, V]) Set(key K, val V) {
	c.cache.Set(key, val)
}

func (c Cache[K, V]) SetWithExp(key K, val V, exp time.Duration) {
	c.cache.Set(key, val, genericsCache.WithExpiration(exp))
}

func (c Cache[K, V]) Delete(key K) {
	c.cache.Delete(key)
}

// NewCache returns a new cache with the given types
func NewCache[K comparable, V any]() v2.Cache[K, V] {
	c := genericsCache.New[K, V]()
	return &Cache[K, V]{
		c,
	}
}
