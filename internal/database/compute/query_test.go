package compute

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuery(t *testing.T) {
	query := NewQuery(GetCommandID, []string{"key", "value"})
	assert.Equal(t, GetCommandID, query.CommandID())
	assert.Equal(t, []string{"key", "value"}, query.Arguments())
}
