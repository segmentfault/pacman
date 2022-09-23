package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

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
func (r *Cache) GetString(ctx context.Context, key string) (string, error) {
	return r.Client.Get(ctx, key).Result()
}

// SetString set string value with key and ttl
func (r *Cache) SetString(ctx context.Context, key, value string, ttl time.Duration) error {
	return r.Client.Set(ctx, key, value, ttl).Err()
}

// GetInt64 get int64 value by key
func (r *Cache) GetInt64(ctx context.Context, key string) (int64, error) {
	return r.Client.Get(ctx, key).Int64()
}

// SetInt64 set int64 value with key and ttl
func (r *Cache) SetInt64(ctx context.Context, key string, value int64, ttl time.Duration) error {
	return r.Client.Set(ctx, key, value, ttl).Err()
}

// Del delete key from cache
func (r *Cache) Del(ctx context.Context, key string) error {
	return r.Client.Del(ctx, key).Err()
}
