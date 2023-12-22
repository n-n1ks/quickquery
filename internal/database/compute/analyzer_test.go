package compute

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAnalyzeQuery(t *testing.T) {
	testTable := []struct {
		tokens []string
		query  Query
		err    error
	}{
		{
			tokens: []string{"GET", "key"},
			query: Query{
				commandID: GetCommandID,
				arguments: []string{"key"},
			},
			err: nil,
		},
		{
			tokens: []string{"SET", "key", "value"},
			query: Query{
				commandID: SetCommandID,
				arguments: []string{"key", "value"},
			},
			err: nil,
		},
		{
			tokens: []string{"DEL", "key"},
			query: Query{
				commandID: DelCommandID,
				arguments: []string{"key"},
			},
			err: nil,
		},
		{
			tokens: []string{"GET"},
			query:  Query{},
			err:    errInvalidCommandArguments,
		},
		{
			tokens: []string{"SET", "foo"},
			query:  Query{},
			err:    errInvalidCommandArguments,
		},
		{
			tokens: []string{"DEL"},
			query:  Query{},
			err:    errInvalidCommandArguments,
		},
	}

	analyzer := NewAnalyzer()
	for _, tt := range testTable {
		t.Run(fmt.Sprintf("tokens: %v", tt.tokens), func(t *testing.T) {
			query, err := analyzer.AnalyzeQuery(tt.tokens)

			assert.Equal(t, tt.query, query)
			require.ErrorIs(t, tt.err, err)
		})
	}
}
