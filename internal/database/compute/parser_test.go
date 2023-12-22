package compute

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseQuery(t *testing.T) {
	testTable := []struct {
		query  string
		tokens []string
		err    error
	}{
		{
			query:  "SET foo bar",
			tokens: []string{"SET", "foo", "bar"},
			err:    nil,
		},
		{
			query:  "SET foo bar\n",
			tokens: []string{"SET", "foo", "bar"},
			err:    nil,
		},
		{
			query:  "   \n   SET foo bar\n\n",
			tokens: []string{"SET", "foo", "bar"},
			err:    nil,
		},
		{
			query:  "SET foo !@#",
			tokens: []string{},
			err:    errInvalidSymbol,
		},
		{
			query:  "GET foo",
			tokens: []string{"GET", "foo"},
			err:    nil,
		},
		{
			query:  "GET foo\n",
			tokens: []string{"GET", "foo"},
			err:    nil,
		},
		{
			query:  "   \n   GET foo\n\n",
			tokens: []string{"GET", "foo"},
			err:    nil,
		},
		{
			query:  "DEL foo",
			tokens: []string{"DEL", "foo"},
			err:    nil,
		},
		{
			query:  "DEL foo\n",
			tokens: []string{"DEL", "foo"},
			err:    nil,
		},
		{
			query:  "   \n   DEL foo\n\r\t   \n",
			tokens: []string{"DEL", "foo"},
			err:    nil,
		},
	}

	for _, tt := range testTable {
		t.Run(fmt.Sprintf("Parse(%s)", tt.query), func(t *testing.T) {
			p := NewParser()
			tokens, err := p.ParseQuery(tt.query)

			assert.ElementsMatch(t, tt.tokens, tokens)
			require.ErrorIs(t, tt.err, err)
		})
	}
}
