package ds

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestQueue(t *testing.T) {
	var q = NewQueue()
	assert.Equal(t, q.Len(), 0)
	q.Push("a")
	q.Push("b")
	assert.Equal(t, q.Len(), 2)
	assert.Equal(t, q.Pop(), "a")
	assert.Equal(t, q.Len(), 1)
}
