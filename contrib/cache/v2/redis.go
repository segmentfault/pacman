package v2

import (
	"time"

	v2 "github.com/segmentfault/pacman/cache/v2"
)

type RedisCacheConfig struct {
	Addr     string `yaml:"addr"`
	UserName string `yaml:"user_name"`
	Password string `yaml:"password"`
	Database uint   `yaml:"database"`
}

// RedisCache is a redis cache wrapper which implemented the cache interface
type RedisCache[K comparable, V any] struct {
}

func (r *RedisCache[K, V]) Get(key K) (value V, ok bool) {
	//TODO implement me
	panic("implement me")
}

func (r *RedisCache[K, V]) Set(key K, val V) {
	//TODO implement me
	panic("implement me")
}

func (r *RedisCache[K, V]) SetWithExp(key K, val V, exp time.Duration) {
	//TODO implement me
	panic("implement me")
}

func (r *RedisCache[K, V]) Delete(key K) {
	//TODO implement me
	panic("implement me")
}

// NewRedisCache returns a new redis cache interface with redis instance configured
func NewRedisCache[K comparable, V any](config RedisCacheConfig) v2.Cache[K, V] {
	return &RedisCache[K, V]{}
}
