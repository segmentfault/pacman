package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/segmentfault/pacman/cache"
)

var _ cache.Cache = (*Cache)(nil)

// Cache redis cache
type Cache struct {
	Client *redis.Client
}

// NewCache creates a new redis cache
func NewCache(conn, username, password string) *Cache {
	client := redis.NewClient(&redis.Options{
		Addr:     conn,
		Username: username,
		Password: password,
	})
	return &Cache{Client: client}
}

// GetString get string value by key
func (r *Cache) GetString(ctx context.Context, key string) (data string, exist bool, err error) {
	data, err = r.Client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", false, nil
	}
	if err != nil {
		return "", false, err
	}
	return data, true, err
}

// SetString set string value with key and ttl
func (r *Cache) SetString(ctx context.Context, key, value string, ttl time.Duration) error {
	return r.Client.Set(ctx, key, value, ttl).Err()
}

// GetInt64 get int64 value by key
func (r *Cache) GetInt64(ctx context.Context, key string) (data int64, exist bool, err error) {
	data, err = r.Client.Get(ctx, key).Int64()
	if err == redis.Nil {
		return 0, false, nil
	}
	if err != nil {
		return 0, false, err
	}
	return data, true, nil
}

// SetInt64 set int64 value with key and ttl
func (r *Cache) SetInt64(ctx context.Context, key string, value int64, ttl time.Duration) error {
	return r.Client.Set(ctx, key, value, ttl).Err()
}

// Increase key by value
func (r *Cache) Increase(ctx context.Context, key string, value int64) (data int64, err error) {
	return r.Client.IncrBy(ctx, key, value).Result()
}

// Decrease key by value
func (r *Cache) Decrease(ctx context.Context, key string, value int64) (data int64, err error) {
	return r.Client.DecrBy(ctx, key, value).Result()
}

// Del delete key from cache
func (r *Cache) Del(ctx context.Context, key string) error {
	return r.Client.Del(ctx, key).Err()
}

// Flush delete all cache entries
func (r *Cache) Flush(ctx context.Context) error {
	r.Client.FlushAll(ctx)
	return nil
}
