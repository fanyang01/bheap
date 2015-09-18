package bheap

import (
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPush(t *testing.T) {
	h := New(CompareInt)
	assert.Equal(t, 0, h.Len())
	assert.True(t, h.IsEmpty())

	rand.Seed(time.Now().UnixNano())
	var s []int
	for i := 0; i < 1<<20; i++ {
		n := rand.Int()
		s = append(s, n)
		h.Push(n)
	}
	assert.False(t, h.IsEmpty())
	assert.Equal(t, 1<<20, h.Len())

	sort.Ints(s)
	for i := 1<<20 - 1; i >= 0; i-- {
		v, ok := h.Top()
		assert.True(t, ok)
		assert.Equal(t, s[i], v.(int))

		v, ok = h.Pop()
		assert.True(t, ok)
		assert.Equal(t, s[i], v.(int))
	}
	v, ok := h.Top()
	assert.False(t, ok)
	assert.Nil(t, v)

	v, ok = h.Pop()
	assert.False(t, ok)
	assert.Nil(t, v)

	h.Clean()
	assert.True(t, h.IsEmpty())

	h2 := New(CompareInt)
	for i := 0; i < 1<<10; i++ {
		h2.Push(i)
	}
	h2 = h2.Merge(nil)
	h2 = h2.Merge(h)
	assert.Equal(t, 1<<10, h2.Len())

	ss := []string{
		"abc",
		"ab",
		"hello",
	}
	h3 := New(CompareString)
	for _, s := range ss {
		h3.Push(s)
	}
}
