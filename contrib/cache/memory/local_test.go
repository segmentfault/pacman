package memory

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCache_GetInt64(t *testing.T) {
	var key = "test"
	var val int64 = 1
	myCache := NewCache()
	err := myCache.SetInt64(context.TODO(), key, val, -1)
	assert.NoError(t, err)

	got, err := myCache.GetInt64(context.TODO(), key)
	assert.NoError(t, err)
	assert.Equal(t, val, got)
}

func TestCache_GetString(t *testing.T) {
	var key = "test"
	var val = "test"
	myCache := NewCache()
	err := myCache.SetString(context.TODO(), key, val, -1)
	assert.NoError(t, err)

	got, err := myCache.GetString(context.TODO(), key)
	assert.NoError(t, err)
	assert.Equal(t, val, got)
}

func TestCache_Del(t *testing.T) {
	var key = "test"
	var val = "test"
	myCache := NewCache()
	err := myCache.SetString(context.TODO(), key, val, -1)
	assert.NoError(t, err)

	err = myCache.Del(context.TODO(), key)
	assert.NoError(t, err)

	got, err := myCache.GetString(context.TODO(), key)
	assert.Equal(t, "", got)
}

func TestCache_Flush(t *testing.T) {
	var key = "test"
	var val = "test"
	myCache := NewCache()
	err := myCache.SetString(context.TODO(), key, val, -1)
	assert.NoError(t, err)

	err = myCache.Flush(context.TODO())
	assert.NoError(t, err)

	got, err := myCache.GetString(context.TODO(), key)
	assert.Equal(t, "", got)
}
