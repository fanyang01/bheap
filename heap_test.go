package bheap

import "testing"

func compare(i, j interface{}) int {
	x, y := i.(int), j.(int)
	if x < y {
		return 1
	}
	return -1
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
