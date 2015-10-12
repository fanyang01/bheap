package bheap

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSimple(t *testing.T) {
	N := 1 << 5
	h := New(CompareInt)
	for i := 0; i < N; i++ {
		// fmt.Println("push", i)
		h.Push(i)
		// printT(h.root)
	}
	for i := N - 1; i >= 0; i-- {
		v, ok := h.Pop()
		assert.True(t, ok)
		assert.Equal(t, i, v.(int))
	}
}

func printT(child []*node) {
	fmt.Printf("[ ")
	for _, n := range child {
		if n == nil {
			fmt.Printf("nil ")
			continue
		}
		fmt.Printf("%d ", n.v)
	}
	fmt.Println("]")
	for _, n := range child {
		if n == nil {
			continue
		}
		fmt.Printf("%d: ", n.v)
		printT(n.child)
	}
}

func TestPush(t *testing.T) {
	N := 1 << 20
	h := New(CompareInt)
	assert.Equal(t, 0, h.Len())
	assert.True(t, h.IsEmpty())

	rand.Seed(time.Now().UnixNano())
	var s []int
	for i := 0; i < N; i++ {
		n := rand.Int()
		s = append(s, n)
		h.Push(n)
	}
	assert.False(t, h.IsEmpty())
	assert.Equal(t, N, h.Len())

	sort.Ints(s)
	for i := N - 1; i >= 0; i-- {
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
