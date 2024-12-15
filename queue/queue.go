package queue

import (
	"fmt"

	"github.com/stygian91/datastructs-go/bst"
)

type Queue[V any] struct {
	q []V
}

var EmptyQueueError = fmt.Errorf("Trying to pop from an empty queue")

func NewQueue[V any]() Queue[V] {
	return Queue[V]{q: []V{}}
}

func (this *Queue[V]) Push(values ...V) {
	this.q = append(this.q, values...)
}

func (this *Queue[V]) Pop() (V, error) {
	if len(this.q) == 0 {
		return *new(V), EmptyQueueError
	}

	res := this.q[0]
	this.q = this.q[1:]

	return res, nil
}

func (this Queue[V]) Len() int {
	return len(this.q)
}

// PriorityQueue uses uint priority values for queue items
// the priority value is inverted meaning that zero is the highest priority item
// and higher numbers have lower priority
type PriorityQueue[V any] struct {
	q   bst.BST[uint, V]
	len int
}

func NewPriorityQueue[V any]() PriorityQueue[V] {
	return PriorityQueue[V]{q: bst.BST[uint, V]{}, len: 0}
}

func (this *PriorityQueue[V]) Add(priority uint, value V) {
	if this.len == 0 {
		this.q.Root = bst.NewNode(priority, value)
		this.len++
		return
	}

	this.q.Add(priority, value)
	this.len++
}

func (this *PriorityQueue[V]) Remove() (V, error) {
	if this.len == 0 {
		return *new(V), EmptyQueueError
	}

	minPrioNode := this.q.Min()
	removedNode, found := this.q.Remove(minPrioNode.Value)

	if !found {
		return *new(V), EmptyQueueError
	}

	this.len--
	return removedNode.Meta, nil
}

func (this PriorityQueue[V]) Len() int {
	return this.len
}

func (this PriorityQueue[V]) Empty() bool {
	return this.len == 0
}
