package database

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"quickquery/internal/database/compute"
	"quickquery/internal/database/storage"
	"quickquery/internal/database/storage/engine/memengine"
	"quickquery/internal/database/storage/engine/mock_engine"
	"quickquery/internal/initialization"
)

func TestNewDatabase(t *testing.T) {
	logger := initialization.NewLogger("fatal")
	ctrl := gomock.NewController(t)
	engine := mock_engine.NewMockEngine(ctrl)
	store, _ := storage.NewStorage(engine, logger)
	parser := compute.NewParser()
	analyzer := compute.NewAnalyzer()
	compute, _ := compute.NewCompute(parser, analyzer, logger)

	t.Run("when compute layer is nil", func(t *testing.T) {
		_, err := NewDatabase(nil, store, logger)

		require.ErrorIs(t, err, errInvalidComputeLayer)
	})

	t.Run("when storage layer is nil", func(t *testing.T) {
		_, err := NewDatabase(compute, nil, logger)

		require.ErrorIs(t, err, errInvalidStorageLayer)
	})

	t.Run("when compute and storage layers are valid", func(t *testing.T) {
		_, err := NewDatabase(compute, store, logger)

		require.NoError(t, err)
	})
}

func TestHandleQuery(t *testing.T) {
	testTable := []struct {
		name    string
		queries []string
		results []string
	}{
		{
			name:    "when command is SET",
			queries: []string{"SET a 10"},
			results: []string{"OK"},
		},
		{
			name:    "when command is SET and GET",
			queries: []string{"SET a 10", "GET a"},
			results: []string{"OK", "10"},
		},
		{
			name:    "when command is GET without value",
			queries: []string{"GET a"},
			results: []string{"ERROR: key not found"},
		},
		{
			name:    "when command is DEL",
			queries: []string{"DEL a"},
			results: []string{"OK"},
		},
	}

	for _, tt := range testTable {
		logger := initialization.NewLogger("fatal")
		engine := memengine.NewEngine()
		store, _ := storage.NewStorage(engine, logger)
		parser := compute.NewParser()
		analyzer := compute.NewAnalyzer()
		compute, _ := compute.NewCompute(parser, analyzer, logger)
		db, _ := NewDatabase(compute, store, logger)

		t.Run(tt.name, func(t *testing.T) {
			for i, query := range tt.queries {
				result := db.HandleQuery(context.TODO(), query)

				assert.Equal(t, tt.results[i], result)
			}
		})
	}
}
