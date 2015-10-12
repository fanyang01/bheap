package bheap

import (
	"container/heap"
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

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

func BenchmarkBHeapPush(b *testing.B) {
	h := New(CompareInt)
	rand.Seed(time.Now().UnixNano())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n := rand.Int()
		h.Push(n)
	}
}
func BenchmarkBHeapPop(b *testing.B) {
	h := New(CompareInt)
	rand.Seed(time.Now().UnixNano())
	var s []int
	for i := 0; i < b.N; i++ {
		n := rand.Int()
		s = append(s, n)
		h.Push(n)
	}
	sort.Ints(s)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h.Pop()
	}
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func BenchmarkStdHeapPush(b *testing.B) {
	h := &IntHeap{}
	heap.Init(h)
	rand.Seed(time.Now().UnixNano())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n := rand.Int()
		heap.Push(h, n)
	}
}

func BenchmarkStdHeapPop(b *testing.B) {
	h := &IntHeap{}
	heap.Init(h)
	rand.Seed(time.Now().UnixNano())
	var s []int
	for i := 0; i < b.N; i++ {
		n := rand.Int()
		heap.Push(h, n)
	}
	sort.Ints(s)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		heap.Pop(h)
	}
}
