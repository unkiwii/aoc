package heap

import (
	"cmp"
	"container/heap"
)

func Less[T cmp.Ordered](a, b T) bool {
	return a < b
}

type Heap[T any] struct {
	data []T
	less func(a, b T) bool
}

func New[T cmp.Ordered]() *Heap[T] {
	h := &Heap[T]{less: Less[T]}
	heap.Init(h)
	return h
}

func NewWithLess[T any](less func(a, b T) bool) *Heap[T] {
	h := &Heap[T]{less: less}
	heap.Init(h)
	return h
}

func (h *Heap[T]) Len() int {
	return len(h.data)
}

// DEPRECATED: DO NOT USE
func (h *Heap[T]) Less(i, j int) bool {
	return h.less(h.data[i], h.data[j])
}

// DEPRECATED: DO NOT USE
func (h Heap[T]) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

// DEPRECATED: Use PushItem instead
func (h *Heap[T]) Push(x any) {
	h.data = append(h.data, x.(T))
}

// DEPRECATED: Use PopItem instead
func (h *Heap[T]) Pop() any {
	old := h.data
	n := len(old)
	x := old[n-1]
	h.data = old[0 : n-1]
	return x
}

func (h *Heap[T]) PushItem(x T) {
	heap.Push(h, x)
}

func (h *Heap[T]) PopItem() T {
	return heap.Pop(h).(T)
}

func (h *Heap[T]) Peek() T {
	if len(h.data) == 0 {
		var zero T
		return zero
	}
	return h.data[0]
}

func (h *Heap[T]) IsEmpty() bool {
	return len(h.data) == 0
}

func (h *Heap[T]) Clear() {
	h.data = h.data[:0]
}

func (h *Heap[T]) Slice() []T {
	r := make([]T, len(h.data))
	copy(r, h.data)
	return r
}
