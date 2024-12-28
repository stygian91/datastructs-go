package heap

import (
	h "container/heap"

	"golang.org/x/exp/constraints"
)

type heap[T constraints.Ordered] []T

func (this heap[T]) Len() int           { return len(this) }
func (this heap[T]) Less(i, j int) bool { return this[i] < this[j] }
func (this heap[T]) Swap(i, j int)      { this[i], this[j] = this[j], this[i] }

func (this *heap[T]) Push(x any) {
	*this = append(*this, x.(T))
}

func (this *heap[T]) Pop() any {
	old := *this
	n := len(old)
	x := old[n-1]
	*this = old[0 : n-1]
	return x
}

type Heap[T constraints.Ordered] struct {
	inner heap[T]
}

func New[T constraints.Ordered]() Heap[T] {
	v := Heap[T]{inner: heap[T]{}}
	h.Init(&v.inner)
	return v
}

func (this *Heap[T]) Push(x any) {
	h.Push(&this.inner, x)
}

func (this *Heap[T]) Pop() any {
	return h.Pop(&this.inner)
}

func (this Heap[T]) Len() int {
	return this.inner.Len()
}
