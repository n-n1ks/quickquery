package compute

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"quickquery/internal/initialization"
)

func TestNewCompute(t *testing.T) {
	logger := initialization.NewLogger("fatal")

	t.Run("when parser is nil", func(t *testing.T) {
		_, err := NewCompute(nil, NewAnalyzer(), logger)

		require.ErrorIs(t, errInvalidParser, err)
	})

	t.Run("when analyzer is nil", func(t *testing.T) {
		_, err := NewCompute(NewParser(), nil, logger)

		require.ErrorIs(t, errInvalidAnalyzer, err)
	})

	t.Run("when parser and analyzer are present", func(t *testing.T) {
		_, err := NewCompute(NewParser(), NewAnalyzer(), logger)

		require.NoError(t, err)
	})
}

func TestHandleQuery(t *testing.T) {
	testTable := []struct {
		queryStr string
		query    Query
		err      error
	}{
		{
			queryStr: "GET key",
			query: Query{
				commandID: GetCommandID,
				arguments: []string{"key"},
			},
			err: nil,
		},
		{
			queryStr: "SET key value",
			query: Query{
				commandID: SetCommandID,
				arguments: []string{"key", "value"},
			},
			err: nil,
		},
		{
			queryStr: "DEL key",
			query: Query{
				commandID: DelCommandID,
				arguments: []string{"key"},
			},
			err: nil,
		},
	}

	logger := initialization.NewLogger("fatal")
	compute, _ := NewCompute(NewParser(), NewAnalyzer(), logger)
	for _, tt := range testTable {
		t.Run(tt.queryStr, func(t *testing.T) {
			query, err := compute.HandleQuery(context.Background(), tt.queryStr)

			assert.Equal(t, tt.query, query)
			require.ErrorIs(t, tt.err, err)
		})
	}
}
