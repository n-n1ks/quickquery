package database

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"quickquery/internal/database/compute"
	database_mock "quickquery/internal/database/mock"
	"quickquery/internal/initialization"
)

func TestNewDatabase(t *testing.T) {
	t.Parallel()

	logger := initialization.NewLogger("fatal")
	ctrl := gomock.NewController(t)
	compute := database_mock.NewMockComputerLayer(ctrl)
	store := database_mock.NewMockStorageLayer(ctrl)

	t.Run("when compute layer is nil", func(t *testing.T) {
		t.Parallel()

		_, err := NewDatabase(nil, store, logger)
		require.ErrorIs(t, err, errInvalidComputeLayer)
	})

	t.Run("when storage layer is nil", func(t *testing.T) {
		t.Parallel()

		_, err := NewDatabase(compute, nil, logger)
		require.ErrorIs(t, err, errInvalidStorageLayer)
	})

	t.Run("when logger is nil", func(t *testing.T) {
		t.Parallel()

		_, err := NewDatabase(compute, store, nil)
		require.ErrorIs(t, err, errInvalidLogger)
	})

	t.Run("when compute and storage layers are valid", func(t *testing.T) {
		t.Parallel()

		_, err := NewDatabase(compute, store, logger)
		require.NoError(t, err)
	})
}

func TestHandleQuery(t *testing.T) {
	t.Parallel()

	t.Run("when command is SET a 10", func(t *testing.T) {
		t.Parallel()

		logger := initialization.NewLogger("fatal")
		ctrl := gomock.NewController(t)

		comp := database_mock.NewMockComputerLayer(ctrl)
		comp.EXPECT().HandleQuery(gomock.Any(), "SET a 10").Return(
			compute.NewQuery(compute.SetCommandID, []string{"a", "10"}), nil,
		)

		store := database_mock.NewMockStorageLayer(ctrl)
		store.EXPECT().Set(gomock.Any(), "a", "10").Return(nil)

		db, _ := NewDatabase(comp, store, logger)
		result := db.HandleQuery(context.TODO(), "SET a 10")

		require.Equal(t, "OK", result)
	})

	t.Run("when command is GET a", func(t *testing.T) {
		t.Parallel()

		logger := initialization.NewLogger("fatal")
		ctrl := gomock.NewController(t)

		comp := database_mock.NewMockComputerLayer(ctrl)
		comp.EXPECT().HandleQuery(gomock.Any(), "GET a").Return(
			compute.NewQuery(compute.GetCommandID, []string{"a"}), nil,
		)

		store := database_mock.NewMockStorageLayer(ctrl)
		store.EXPECT().Get(gomock.Any(), "a").Return("10", nil)

		db, _ := NewDatabase(comp, store, logger)
		result := db.HandleQuery(context.TODO(), "GET a")

		require.Equal(t, "10", result)
	})

	t.Run("when command is DEL a", func(t *testing.T) {
		t.Parallel()

		logger := initialization.NewLogger("fatal")
		ctrl := gomock.NewController(t)

		comp := database_mock.NewMockComputerLayer(ctrl)
		comp.EXPECT().HandleQuery(gomock.Any(), "DEL a").Return(
			compute.NewQuery(compute.DelCommandID, []string{"a"}), nil,
		)

		store := database_mock.NewMockStorageLayer(ctrl)
		store.EXPECT().Del(gomock.Any(), "a").Return(nil)

		db, _ := NewDatabase(comp, store, logger)
		result := db.HandleQuery(context.TODO(), "DEL a")

		require.Equal(t, "OK", result)
	})
}
