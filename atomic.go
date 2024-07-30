package gtl

import (
	"sync/atomic"
	"unsafe"
)

type AtomicPointer[T any] struct {
	ptr unsafe.Pointer
}

func (r *AtomicPointer[T]) Load() *T {
	return (*T)(atomic.LoadPointer(&r.ptr))
}

func (r *AtomicPointer[T]) Store(val *T) {
	atomic.StorePointer(&r.ptr, unsafe.Pointer(val))
}
