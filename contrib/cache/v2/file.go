package v2

import (
	"time"

	v2 "github.com/segmentfault/pacman/cache/v2"
)

type FilesCache[K comparable, V any] struct {
}

func (f *FilesCache[K, V]) Get(key K) (value V, ok bool) {
	//TODO implement me
	panic("implement me")
}

func (f *FilesCache[K, V]) Set(key K, val V) {
	//TODO implement me
	panic("implement me")
}

func (f *FilesCache[K, V]) SetWithExp(key K, val V, exp time.Duration) {
	//TODO implement me
	panic("implement me")
}

func (f *FilesCache[K, V]) Delete(key K) {
	//TODO implement me
	panic("implement me")
}

// NewFilesCache returns cache instance for local file system storage
func NewFilesCache[K comparable, V any](dir string) v2.Cache[K, V] {
	return &FilesCache[K, V]{}
}
