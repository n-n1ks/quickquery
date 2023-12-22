package memengine

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	engine := NewEngine()

	key := "key"
	value := "value"

	engine.Set(context.Background(), key, value)
	result, ok := engine.Get(context.Background(), key)

	assert.Equal(t, value, result)
	assert.True(t, ok)
}

func TestDel(t *testing.T) {
	engine := NewEngine()

	key := "key"
	value := "value"

	engine.Set(context.Background(), key, value)
	result, ok := engine.Get(context.Background(), key)

	assert.Equal(t, value, result)
	assert.True(t, ok)

	engine.Del(context.Background(), key)

	result, ok = engine.Get(context.Background(), key)

	assert.Equal(t, "", result)
	assert.False(t, ok)
}
