package v2

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type Person struct {
	Name string
}

func TestNewCache(t *testing.T) {
	cache := NewCache[int, string]()
	cache.Set(1, "hello")

	got, ok := cache.Get(1)
	assert.Equal(t, ok, true)
	assert.Equal(t, got, "hello")
}

func TestNewPersonCache(t *testing.T) {
	cache := NewCache[string, Person]()
	cache.Set("john", Person{Name: "John Smith"})

	got, ok := cache.Get("john")
	assert.Equal(t, ok, true)
	assert.NotNil(t, got)
	assert.Equal(t, got.Name, "John Smith")
}

func TestExpireCache(t *testing.T) {
	cache := NewCache[string, Person]()
	cache.SetWithExp("john", Person{Name: "John Smith"}, time.Second)

	got, ok := cache.Get("john")
	assert.Equal(t, ok, true)
	assert.NotNil(t, got)
	assert.Equal(t, got.Name, "John Smith")

	time.Sleep(2 * time.Second)
	got, ok = cache.Get("john")
	assert.NotEqual(t, ok, true)
	assert.Empty(t, got)
}
