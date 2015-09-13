// Package bheap implements binomial-heap.
package bheap

type heapTree struct {
	siblings *heapTree
	childs   *heapTree
	degree   int
	data     interface{}
}

// Heap is a binomial-heap.
type Heap struct {
	list    *heapTree
	size    int
	compare Comparator
}

// Comparator compares x and y, and returns an integer
// = 0 if x is equal to y,
// > 0 if x is greater than y, and
// < 0 if x is less than y
type Comparator func(x, y interface{}) int

// New returns an initialized heap.
func New(cmp Comparator) *Heap {
	return &Heap{
		list:    nil,
		size:    0,
		compare: cmp,
	}
}

// IsEmpty returns true if h is empty, otherwise false.
func (h *Heap) IsEmpty() bool {
	return h.size == 0
}

// Clean cleans a heap and sets it to initial state.
func (h *Heap) Clean() *Heap {
	h.size = 0
	h.list = nil
	return h
}

// Merge merges x into h. Note that x is not preserved.
func (h *Heap) Merge(x *Heap) *Heap {
	if x == nil {
		return h
	}
	h.list = h.merge(h.list, x.list)
	h.size += x.size
	return h
}

// merge is the core function of this data structure.
func (h *Heap) merge(x, y *heapTree) *heapTree {
	if x == nil {
		return y
	}
	if y == nil {
		return x
	}

	if x.degree > y.degree {
		y = h.merge(x.siblings, y)
		x.siblings = nil
		if y.degree < x.degree {
			x.siblings = y
			return x
		}
	}
	if x.degree == y.degree {
		rest := h.merge(x.siblings, y.siblings)
		if h.compare(x.data, y.data) < 0 {
			x, y = y, x
		}
		y.siblings = x.childs
		x.childs = y
		x.degree++
		x.siblings = rest
		return x
	}
	return h.merge(y, x)
}

// Pop pops the element that has the highest priority.
func (h *Heap) Pop() interface{} {
	if h.size == 0 {
		return nil
	}

	highest := h.list
	ptrToHighest := &h.list
	pos := h.list.siblings
	prev := &h.list.siblings

	for {
		if pos == nil {
			break
		}
		if h.compare(highest.data, pos.data) < 0 {
			highest = pos
			ptrToHighest = prev
		}
		prev = &pos.siblings
		pos = pos.siblings
	}

	*ptrToHighest = highest.siblings
	h.list = h.merge(h.list, highest.childs)
	h.size--
	return highest.data
}

// Top returns the element that has the highest priority.
func (h *Heap) Top() interface{} {
	if h.size == 0 {
		return nil
	}

	highest := h.list
	pos := h.list.siblings

	for {
		if pos == nil {
			break
		}
		if h.compare(highest.data, pos.data) < 0 {
			highest = pos
		}
		pos = pos.siblings
	}

	return highest.data
}

// Push inserts x into h.
func (h *Heap) Push(x interface{}) {
	t := newTree(x)
	h.list = h.merge(h.list, t)
	h.size++
}

// helper function
func newTree(x interface{}) *heapTree {
	return &heapTree{
		degree:   0,
		data:     x,
		siblings: nil,
		childs:   nil,
	}
}
