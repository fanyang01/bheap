package bheap

type node struct {
	child []*node
	rank  int
	v     interface{}
}

// Heap is a binomial-heap.
type Heap struct {
	root []*node
	size int
	less LessFunc
}

// New returns an initialized heap.
func New(less LessFunc) *Heap {
	return &Heap{
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
	h.root = nil
	return h
}

// Merge merges x into h. Note that x is not preserved.
func (h *Heap) Merge(x *Heap) *Heap {
	if x == nil {
		return h
	}
	h.root = h.merge(h.root, x.root)
	h.size += x.size
	return h
}

// merge is the core function of this data structure.
func (h *Heap) merge(x, y []*node) []*node {
	if len(x) == 0 {
		return y
	}
	if len(y) == 0 {
		return x
	}

	if x[0].rank > y[0].rank {
		y = h.merge(x[1:], y)
		x = []*node{x[0]}
		if y[0].rank < x[0].rank {
			x = append(x, y...)
			return x
		}
	}
	if x[0].rank == y[0].rank {
		rest := h.merge(x[1:], y[1:])
		x, y = []*node{x[0]}, []*node{y[0]}
		if h.less(x[0].v, y[0].v) {
			x, y = y, x
		}
		y = append(y, x[0].child...)
		x[0].child = y
		x[0].rank++
		x = append(x, rest...)
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
	var i, max int
	for i = 1; i < len(h.root); i++ {
		if h.less(h.root[max].v, h.root[i].v) {
			max = i
		}
	}
	t := h.root[max]
	h.root = append(h.root[:max], h.root[max+1:]...)
	h.root = h.merge(h.root, t.child)
	h.size--
	return t.v, true
}

// Top returns the element that has the highest priority.
func (h *Heap) Top() (v interface{}, ok bool) {
	if h.size == 0 {
		return
	}
	var i, max int
	for i = 1; i < len(h.root); i++ {
		if h.less(h.root[max].v, h.root[i].v) {
			max = i
		}
	}
	return h.root[max].v, true
}

// Push inserts x into h.
func (h *Heap) Push(x interface{}) {
	h.root = h.merge(h.root, []*node{&node{v: x}})
	h.size++
}
