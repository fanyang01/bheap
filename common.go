/*
This package uses callbacks to compare items. Using tricks to get pointer
of empty interface values can avoid data copying and runtime assertions,
therefore greatly improve performance.  It's your responsibility to assure
type safe.
*/
package bheap

import "unsafe"

// LessFunc compares x and y, and returns
// true if x is less than y,
// false if x is not less than y.
type LessFunc func(x, y interface{}) bool

// These functions are provided for convinence
var (
	CompareInt    LessFunc = compareInt
	CompareString          = compareString
)

type iface struct {
	typ  unsafe.Pointer
	data unsafe.Pointer
}

// ValuePtr is a helper function to get the pointer to value stored in empty interface.
func ValuePtr(v interface{}) unsafe.Pointer {
	return ((*iface)(unsafe.Pointer(&v))).data
}

func compareInt(x, y interface{}) bool {
	a := *(*int)(ValuePtr(x))
	b := *(*int)(ValuePtr(y))
	return a < b
}

func compareString(x, y interface{}) bool {
	a := *(*string)(ValuePtr(x))
	b := *(*string)(ValuePtr(y))
	return a < b
}
