package memengine

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGet(t *testing.T) {
	t.Parallel()

	t.Run("when key exists", func(t *testing.T) {
		t.Parallel()

		engine := NewEngine()

		key := "key"
		value := "value"

		engine.Set(context.Background(), key, value)
		result, ok := engine.Get(context.Background(), key)

		require.Equal(t, value, result)
		require.True(t, ok)
	})

	t.Run("when key doesn't exist", func(t *testing.T) {
		t.Parallel()

		engine := NewEngine()

		key := "key"

		result, ok := engine.Get(context.Background(), key)

		require.Equal(t, "", result)
		require.False(t, ok)
	})
}

func TestSet(t *testing.T) {
	t.Parallel()

	engine := NewEngine()

	key := "key"
	value := "value"

	engine.Set(context.Background(), key, value)
	result, ok := engine.Get(context.Background(), key)

	require.Equal(t, value, result)
	require.True(t, ok)
}

func TestDel(t *testing.T) {
	t.Parallel()
	engine := NewEngine()

	key := "key"
	value := "value"

	engine.Set(context.Background(), key, value)
	result, ok := engine.Get(context.Background(), key)

	require.Equal(t, value, result)
	require.True(t, ok)

	engine.Del(context.Background(), key)

	result, ok = engine.Get(context.Background(), key)

	require.Equal(t, "", result)
	require.False(t, ok)
}
