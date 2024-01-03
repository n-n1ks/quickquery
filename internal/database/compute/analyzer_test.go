package compute

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAnalyzeQuery(t *testing.T) {
	t.Parallel()

	testTable := []struct {
		name   string
		tokens []string
		query  Query
		err    error
	}{
		{
			name:   "when command is GET with a key",
			tokens: []string{"GET", "key"},
			query: Query{
				commandID: GetCommandID,
				arguments: []string{"key"},
			},
			err: nil,
		},
		{
			name:   "when command is SET with a key and value",
			tokens: []string{"SET", "key", "value"},
			query: Query{
				commandID: SetCommandID,
				arguments: []string{"key", "value"},
			},
			err: nil,
		},
		{
			name:   "when command is DEL with a key",
			tokens: []string{"DEL", "key"},
			query: Query{
				commandID: DelCommandID,
				arguments: []string{"key"},
			},
			err: nil,
		},
		{
			name:   "when command is GET without a key",
			tokens: []string{"GET"},
			query:  Query{},
			err:    errInvalidCommandArguments,
		},
		{
			name:   "when command is SET with a key but without value",
			tokens: []string{"SET", "foo"},
			query:  Query{},
			err:    errInvalidCommandArguments,
		},
		{
			name:   "when command is DEL without a key",
			tokens: []string{"DEL"},
			query:  Query{},
			err:    errInvalidCommandArguments,
		},
	}

	for _, tt := range testTable {
		test := tt
		testAnalyzer := NewAnalyzer()

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			query, err := testAnalyzer.AnalyzeQuery(test.tokens)

			assert.Equal(t, test.query, query)
			require.ErrorIs(t, test.err, err)
		})
	}
}
