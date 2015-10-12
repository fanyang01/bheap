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
	max  int
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
	h.max = -1
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

func (h *Heap) merge(x, y []*node) []*node {
	if len(y) == 0 {
		return x
	}

	R := len(x) - 1
	if R < len(y)-1 {
		R = len(y) - 1
	}
	for i := len(x) - 1; i < R; i++ {
		x = append(x, nil)
	}
	var carry *node
	for i, j := 0, 0; i <= R; i++ {
		flag := 0
		if x[i] != nil && x[i].rank == i {
			flag |= 01
		}
		if j < len(y) && y[j] != nil && y[j].rank == i {
			flag |= 02
		}
		if carry != nil {
			flag |= 04
		}
		switch flag {
		case 0:
		case 1:
		case 2:
			x[i] = y[j]
			j++
		case 4:
			x[i], carry = carry, nil
		case 3:
			carry = h.combine(x[i], y[j])
			x[i] = nil
			j++
		case 5:
			carry = h.combine(x[i], carry)
			x[i] = nil
		case 6:
			carry = h.combine(y[j], carry)
			j++
		case 7:
			t := x[i]
			x[i], carry = carry, h.combine(t, y[j])
			j++
		}
	}
	if carry != nil {
		x = append(x, carry)
	}
	for R = len(x) - 1; x[R] == nil; R-- {
	}
	x = x[:R+1]
	return x
}

// x.rank == y.rank
func (h *Heap) combine(x, y *node) *node {
	if h.less(x.v, y.v) {
		x, y = y, x
	}
	x.child = append(x.child, y)
	x.rank++
	return x
}

func (h *Heap) updateMax() {
	var i, max int
	for i, max = 0, -1; i < len(h.root); i++ {
		if h.root[i] != nil && (max == -1 || h.less(h.root[max].v, h.root[i].v)) {
			max = i
		}
	}
	h.max = max
}

// Pop returns the element that has the highest priority and a boolean value
// which indicates whether the heap is not empty.
func (h *Heap) Pop() (v interface{}, ok bool) {
	if h.size == 0 {
		return
	}
	t := h.root[h.max]
	h.root[h.max] = nil
	h.root = h.merge(h.root, t.child)
	h.size--
	h.updateMax()
	return t.v, true
}

// Top returns the element that has the highest priority.
func (h *Heap) Top() (v interface{}, ok bool) {
	if h.size == 0 {
		return
	}
	return h.root[h.max].v, true
}

// Push inserts x into h.
func (h *Heap) Push(x interface{}) {
	h.root = h.merge(h.root, []*node{&node{v: x}})
	h.size++
	h.updateMax()
}
