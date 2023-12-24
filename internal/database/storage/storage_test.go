package storage

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"quickquery/internal/database/storage/engine/memengine"
	"quickquery/internal/database/storage/engine/mock_engine"
	"quickquery/internal/initialization"
)

func TestNewStorage(t *testing.T) {
	t.Parallel()

	logger := initialization.NewLogger("fatal")

	t.Run("when engine is incorrect", func(t *testing.T) {
		t.Parallel()

		_, err := NewStorage(nil, logger)
		require.ErrorIs(t, errInvalidEngine, err)
	})

	t.Run("when logger is incorrect", func(t *testing.T) {
		t.Parallel()

		_, err := NewStorage(memengine.NewEngine(), nil)
		require.ErrorIs(t, errInvalidLogger, err)
	})

	t.Run("when engine is correct", func(t *testing.T) {
		t.Parallel()

		_, err := NewStorage(memengine.NewEngine(), logger)
		require.NoError(t, err)
	})
}

func TestGet(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	engine := mock_engine.NewMockEngine(ctrl)
	storage := &Storage{
		engine: engine,
	}

	t.Run("when key is not found", func(t *testing.T) {
		t.Parallel()

		key := "key"
		engine.EXPECT().Get(gomock.Any(), key).Return("", false)

		_, err := storage.Get(context.TODO(), key)
		assert.Equal(t, errKeyNotFound, err)
	})

	t.Run("when key is found", func(t *testing.T) {
		t.Parallel()

		key := "key"
		value := "value"
		engine.EXPECT().Get(gomock.Any(), key).Return(value, true)

		result, err := storage.Get(context.TODO(), key)
		require.NoError(t, err)
		assert.Equal(t, value, result)
	})
}

func TestSet(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	engine := mock_engine.NewMockEngine(ctrl)
	storage := &Storage{
		engine: engine,
	}

	key := "key"
	engine.EXPECT().Set(gomock.Any(), key, gomock.Any())

	err := storage.Set(context.TODO(), key, "value")

	require.NoError(t, err)
}

func TestDel(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	engine := mock_engine.NewMockEngine(ctrl)
	storage := &Storage{
		engine: engine,
	}

	key := "key"
	engine.EXPECT().Del(gomock.Any(), key)

	_ = storage.Del(context.TODO(), key)
}
