package set

import (
	"iter"
	"maps"
)

type empty struct{}

type Set[K comparable] struct {
	m map[K]empty
}

func New[K comparable]() Set[K] {
	return Set[K]{m: map[K]empty{}}
}

func (this *Set[K]) Add(el K) bool {
	if this.Contains(el) {
		return true
	}

	this.m[el] = empty{}
	return false
}

func (this *Set[K]) Remove(el K) bool {
	if !this.Contains(el) {
		return false
	}

	delete(this.m, el)
	return true
}

func (this Set[K]) Contains(el K) bool {
	_, exists := this.m[el]
	return exists
}

func (this Set[K]) Len() int {
	return len(this.m)
}

func (this Set[K]) Seq() iter.Seq[K] {
	return maps.Keys(this.m)
}

func (this *Set[K]) Merge(other Set[K]) {
	for v := range other.Seq() {
		this.m[v] = empty{}
	}
}
