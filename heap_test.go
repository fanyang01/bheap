package bheap

import "testing"

type Int int

func (i Int) Less(j Interface) bool {
	if i > j.(Int) {
		return true
	}
	return false
}

func TestPush(t *testing.T) {
	h := New()
	for i := 0; i < (1 << 20); i++ {
		h.Push(Int(i))
	}
	for i := 0; i < (1 << 20); i++ {
		if h.Pop().(Int) != Int(i) {
			t.Fail()
		}
	}
}
