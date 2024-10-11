package stack

import "fmt"

type Stack[V any] struct {
	s []V
}

var EmptyStackError = fmt.Errorf("Trying to pop an empty stack.")

func NewStack[V any]() Stack[V] {
	return Stack[V]{}
}

func (this *Stack[V]) Push(vals ...V) {
	this.s = append(this.s, vals...)
}

func (this *Stack[V]) Pop() (V, error) {
	if len(this.s) == 0 {
		return *new(V), EmptyStackError
	}

	res := this.s[len(this.s)-1]
	this.s = this.s[:len(this.s)-1]

	return res, nil
}

func (this Stack[V]) Len() int {
	return len(this.s)
}
