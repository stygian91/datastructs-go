package queue

import "fmt"

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
