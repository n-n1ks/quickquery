package compute

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"quickquery/internal/initialization"
)

func TestNewCompute(t *testing.T) {
	t.Parallel()

	logger := initialization.NewLogger("fatal")

	t.Run("when parser is nil", func(t *testing.T) {
		t.Parallel()

		_, err := NewCompute(nil, NewAnalyzer(), logger)
		require.ErrorIs(t, errInvalidParser, err)
	})

	t.Run("when analyzer is nil", func(t *testing.T) {
		t.Parallel()

		_, err := NewCompute(NewParser(), nil, logger)
		require.ErrorIs(t, errInvalidAnalyzer, err)
	})

	t.Run("when logger is nil", func(t *testing.T) {
		t.Parallel()

		_, err := NewCompute(NewParser(), NewAnalyzer(), nil)
		require.ErrorIs(t, errInvalidLogger, err)
	})

	t.Run("when parser and analyzer are present", func(t *testing.T) {
		t.Parallel()

		_, err := NewCompute(NewParser(), NewAnalyzer(), logger)
		require.NoError(t, err)
	})
}

func TestHandleQuery(t *testing.T) {
	t.Parallel()

	testTable := []struct {
		name     string
		queryStr string
		query    Query
		err      error
	}{
		{
			name:     "when command is GET with a key",
			queryStr: "GET key",
			query: Query{
				commandID: GetCommandID,
				arguments: []string{"key"},
			},
			err: nil,
		},
		{
			name:     "when command is SET with a key and value",
			queryStr: "SET key value",
			query: Query{
				commandID: SetCommandID,
				arguments: []string{"key", "value"},
			},
			err: nil,
		},
		{
			name:     "when command is DEL with a key",
			queryStr: "DEL key",
			query: Query{
				commandID: DelCommandID,
				arguments: []string{"key"},
			},
			err: nil,
		},
	}

	for _, tt := range testTable {
		test := tt
		logger := initialization.NewLogger("fatal")
		compute, _ := NewCompute(NewParser(), NewAnalyzer(), logger)

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			query, err := compute.HandleQuery(context.Background(), test.queryStr)

			assert.Equal(t, test.query, query)
			require.ErrorIs(t, test.err, err)
		})
	}
}
