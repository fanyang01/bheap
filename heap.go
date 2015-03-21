package bheap

// Interface is the data store in heap, must be comparable
type Interface interface {
	// Less compare the priority of reciever with argument
	Less(Interface) bool
}

type heapTree struct {
	siblings *heapTree
	childs   *heapTree
	degree   int
	data     Interface
}

// Heap is the head of the structure
type Heap struct {
	list *heapTree
	size int
}

// New return a initialized heap
func New() *Heap {
	return &Heap{
		list: nil,
		size: 0,
	}
}

// IsEmpty return true if heap is empty, otherwise false
func (h *Heap) IsEmpty() bool {
	return h.size == 0
}

// Clean clean a heap, set it to initial stat
func (h *Heap) Clean() *Heap {
	h.size = 0
	h.list = nil
	return h
}

// Merge merge heap x into heap h
func (h *Heap) Merge(x *Heap) *Heap {
	if x == nil {
		return h
	}
	h.list = merge(h.list, x.list)
	h.size += x.size
	return h
}

// Pop pop the element with highest priority
func (h *Heap) Pop() Interface {
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
		if highest.data.Less(pos.data) {
			highest = pos
			ptrToHighest = prev
		}
		prev = &pos.siblings
		pos = pos.siblings
	}

	*ptrToHighest = highest.siblings
	merge(h.list, highest.childs)
	h.size--
	return highest.data
}

// Top return the element with highest priority
func (h *Heap) Top() Interface {
	if h.size == 0 {
		return nil
	}

	highest := h.list
	pos := h.list.siblings

	for {
		if pos == nil {
			break
		}
		if highest.data.Less(pos.data) {
			highest = pos
		}
		pos = pos.siblings
	}

	return highest.data
}

// Push push data x which implement Interface into heap
func (h *Heap) Push(x Interface) {
	t := newTree(x)
	h.list = merge(h.list, t)
	h.size++
}

// merge is the core of this data structure
func merge(x, y *heapTree) *heapTree {
	if x == nil {
		return y
	}
	if y == nil {
		return x
	}

	if x.degree > y.degree {
		y = merge(x.siblings, y)
		x.siblings = nil
		if y.degree < x.degree {
			x.siblings = y
			return x
		}
	}
	if x.degree == y.degree {
		rest := merge(x.siblings, y.siblings)
		if x.data.Less(y.data) {
			x, y = y, x
		}
		y.siblings = x.childs
		x.childs = y
		x.degree++
		x.siblings = rest
		return x
	}
	return merge(y, x)
}

func newTree(x Interface) *heapTree {
	return &heapTree{
		degree:   0,
		data:     x,
		siblings: nil,
		childs:   nil,
	}
}
