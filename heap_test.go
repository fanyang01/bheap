package bheap

import "testing"

func compare(x, y interface{}) bool {
	return !CompareInt(x, y)
}

func TestPush(t *testing.T) {
	h := New(compare)
	for i := 0; i < 1<<20; i++ {
		h.Push(i)
	}
	for i := 0; i < 1<<20; i++ {
		if h.Pop().(int) != i {
			t.Fail()
		}
	}
}
