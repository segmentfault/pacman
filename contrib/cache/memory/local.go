package memory

import (
	"context"
	"fmt"
	"strconv"
	"time"

	goCache "github.com/patrickmn/go-cache"
)

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
func (s *Cache) GetString(ctx context.Context, key string) (string, error) {
	value, has := s.local.Get(key)
	if !has {
		return "", fmt.Errorf("information does not exist")
	}
	v, ok := value.(string)
	if !ok {
		return "", fmt.Errorf("information abnormality")
	}
	return v, nil
}

// SetString set string value with key and ttl
func (s *Cache) SetString(ctx context.Context, key, value string, ttl time.Duration) error {
	s.local.Set(key, value, ttl)
	return nil
}

// GetInt64 get int64 value by key
func (s *Cache) GetInt64(ctx context.Context, key string) (int64, error) {
	value, has := s.local.Get(key)
	if !has {
		return 0, fmt.Errorf("information does not exist")
	}
	v, ok := value.(string)
	if !ok {
		return 0, fmt.Errorf("information abnormality")
	}
	num, _ := strconv.ParseInt(v, 10, 64)
	return num, nil
}

// SetInt64 set int64 value with key and ttl
func (s *Cache) SetInt64(ctx context.Context, key string, value int64, ttl time.Duration) error {
	str := fmt.Sprintf("%d", value)
	s.local.Set(key, str, ttl)
	return nil
}

// Del delete key from cache
func (s *Cache) Del(ctx context.Context, Key string) error {
	s.local.Delete(Key)
	return nil
}
