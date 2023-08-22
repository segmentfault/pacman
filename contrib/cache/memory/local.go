package memory

import (
	"context"
	"fmt"
	"time"

	goCache "github.com/patrickmn/go-cache"
	"github.com/segmentfault/pacman/cache"
)

var _ cache.Cache = (*Cache)(nil)

// Cache basic memory cache
type Cache struct {
	local *goCache.Cache
}

// NewCache new memory cache
func NewCache() *Cache {
	return &Cache{local: goCache.New(goCache.NoExpiration, 10*time.Minute)}
}

// Save sync cache file to disk file, just for testing
func Save(memCache *Cache, filePath string) error {
	return memCache.local.SaveFile(filePath)
}

// Load read cache file from disk file, just for testing
func Load(memCache *Cache, filePath string) error {
	return memCache.local.LoadFile(filePath)
}

// GetString get string value by key
func (c *Cache) GetString(ctx context.Context, key string) (data string, exist bool, err error) {
	value, has := c.local.Get(key)
	if !has {
		return "", false, nil
	}
	data, _ = value.(string)
	return data, true, nil
}

// SetString set string value with key and ttl
func (c *Cache) SetString(ctx context.Context, key, value string, ttl time.Duration) error {
	c.local.Set(key, value, ttl)
	return nil
}

// GetInt64 get int64 value by key
func (c *Cache) GetInt64(ctx context.Context, key string) (data int64, exist bool, err error) {
	value, has := c.local.Get(key)
	if !has {
		return 0, false, nil
	}
	data, _ = value.(int64)
	return data, true, nil
}

// SetInt64 set int64 value with key and ttl
func (c *Cache) SetInt64(ctx context.Context, key string, value int64, ttl time.Duration) error {
	str := fmt.Sprintf("%d", value)
	c.local.Set(key, str, ttl)
	return nil
}

func (c *Cache) Increase(ctx context.Context, key string, value int64) (data int64, err error) {
	return c.local.IncrementInt64(key, value)
}

func (c *Cache) Decrease(ctx context.Context, key string, value int64) (data int64, err error) {
	return c.local.DecrementInt64(key, value)
}

// Del delete key from cache
func (c *Cache) Del(ctx context.Context, Key string) error {
	c.local.Delete(Key)
	return nil
}

// Flush deletes all items from cache
func (c *Cache) Flush(ctx context.Context) error {
	c.local.Flush()
	return nil
}
