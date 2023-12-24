package compute

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseQuery(t *testing.T) {
	t.Parallel()

	testTable := []struct {
		name   string
		query  string
		tokens []string
		err    error
	}{
		{
			name:   "when a command is SET with a valid key and value",
			query:  "SET foo bar",
			tokens: []string{"SET", "foo", "bar"},
			err:    nil,
		},
		{
			name:   "when a command is SET with a valid key and value and a trailing newline",
			query:  "SET foo bar\n",
			tokens: []string{"SET", "foo", "bar"},
			err:    nil,
		},
		{
			name:   "when a command is SET with a valid key and value and a trailing newline and spaces",
			query:  "   \n   SET foo bar\n\n",
			tokens: []string{"SET", "foo", "bar"},
			err:    nil,
		},
		{
			name:   "when a command is SET with a valid key and invalid value",
			query:  "SET foo !@#",
			tokens: []string{},
			err:    errInvalidSymbol,
		},
		{
			name:   "when a command is GET with a valid key",
			query:  "GET foo",
			tokens: []string{"GET", "foo"},
			err:    nil,
		},
		{
			name:   "when a command is GET with a valid key and a trailing newline",
			query:  "GET foo\n",
			tokens: []string{"GET", "foo"},
			err:    nil,
		},
		{
			name:   "when a command is GET with a valid key and a trailing newline and spaces",
			query:  "   \n   GET foo\n\n",
			tokens: []string{"GET", "foo"},
			err:    nil,
		},
		{
			name:   "when a command is DEL with a valid key",
			query:  "DEL foo",
			tokens: []string{"DEL", "foo"},
			err:    nil,
		},
		{
			name:   "when a command is DEL with a valid key and a trailing newline",
			query:  "DEL foo\n",
			tokens: []string{"DEL", "foo"},
			err:    nil,
		},
		{
			name:   "when a command is DEL with a valid key and a trailing newline and spaces",
			query:  "   \n   DEL foo\n\r\t   \n",
			tokens: []string{"DEL", "foo"},
			err:    nil,
		},
	}

	for _, tt := range testTable {
		test := tt

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			p := NewParser()
			tokens, err := p.ParseQuery(test.query)

			assert.ElementsMatch(t, test.tokens, tokens)
			require.ErrorIs(t, test.err, err)
		})
	}
}
