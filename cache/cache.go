package cache

import (
	"context"
	"time"
)

// Deprecated: use generics cache instead
type Cache interface {
	GetString(ctx context.Context, key string) (string, error)
	SetString(ctx context.Context, key, value string, ttl time.Duration) error
	GetInt64(ctx context.Context, key string) (int64, error)
	SetInt64(ctx context.Context, key string, value int64, ttl time.Duration) error
	Del(ctx context.Context, key string) error
	Flush(ctx context.Context) error
}
