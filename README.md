# Binomial Heap [![GoDoc](https://godoc.org/github.com/fanyang01/bheap?status.svg)](https://godoc.org/github.com/fanyang01/bheap) [![Circle CI](https://circleci.com/gh/fanyang01/bheap.svg?style=svg)](https://circleci.com/gh/fanyang01/bheap) [![Build Status](https://drone.io/github.com/fanyang01/bheap/status.png)](https://drone.io/github.com/fanyang01/bheap/latest)

Package bheap provides binomial-heap written in Go. Unlike the heap package provided by standard library, you don't need to implement any interface.

## Example usage:

```go
	import "github.com/fanyang01/bheap"

	func compare(x, y interface{}) bool {
		return bheap.CompareInt(y, x)
	}

	func test() {
		h := bheap.New(compare)
		for i := 0; i < 1<<20; i++ {
			h.Push(i)
		}
		for i := 0; i < 1<<20; i++ {
			if h.Pop().(int) != i {
				// error
			}
		}
	}
```
