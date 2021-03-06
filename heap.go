package bheap

type heapTree struct {
	siblings *heapTree
	child    *heapTree
	degree   int
	data     interface{}
}

// Heap is a binomial-heap.
type Heap struct {
	list *heapTree
	size int
	less LessFunc
}

// New returns an initialized heap.
func New(less LessFunc) *Heap {
	return &Heap{
		list: nil,
		size: 0,
		less: less,
	}
}

// IsEmpty returns true if h is empty, otherwise false.
func (h *Heap) IsEmpty() bool {
	return h.size == 0
}

// Len returns the number of elements in h.
func (h *Heap) Len() int {
	return h.size
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
		if h.less(x.data, y.data) {
			x, y = y, x
		}
		y.siblings = x.child
		x.child = y
		x.degree++
		x.siblings = rest
		return x
	}
	return h.merge(y, x)
}

// Pop returns the element that has the highest priority and a boolean value
// which indicates whether the heap is not empty.
func (h *Heap) Pop() (v interface{}, ok bool) {
	if h.size == 0 {
		return
	}

	highest := h.list
	ptrToHighest := &h.list
	pos := h.list.siblings
	prev := &h.list.siblings

	for {
		if pos == nil {
			break
		}
		if h.less(highest.data, pos.data) {
			highest = pos
			ptrToHighest = prev
		}
		prev = &pos.siblings
		pos = pos.siblings
	}

	*ptrToHighest = highest.siblings
	h.list = h.merge(h.list, highest.child)
	h.size--
	return highest.data, true
}

// Top returns the element that has the highest priority.
func (h *Heap) Top() (v interface{}, ok bool) {
	if h.size == 0 {
		return
	}

	highest := h.list
	pos := h.list.siblings

	for {
		if pos == nil {
			break
		}
		if h.less(highest.data, pos.data) {
			highest = pos
		}
		pos = pos.siblings
	}

	return highest.data, true
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
		child:    nil,
	}
}
