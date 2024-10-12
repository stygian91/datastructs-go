package stack

import "fmt"

type Stack[V any] struct {
	s []V
}

var EmptyStackError = fmt.Errorf("Trying to pop an empty stack.")

func NewStack[V any]() Stack[V] {
	return Stack[V]{s: []V{}}
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

type SizedStack[V any] struct {
	s      []V
	i, cap int
}

var StackOverflowError = fmt.Errorf("Trying to push to a full stack.")

func NewSizedStack[V any](cap int) SizedStack[V] {
	return SizedStack[V]{i: 0, cap: cap, s: make([]V, cap)}
}

func (this *SizedStack[V]) Push(vals ...V) error {
	for _, val := range vals {
		if this.i >= this.cap {
			return StackOverflowError
		}

		this.s[this.i] = val
		this.i++
	}

	return nil
}

func (this *SizedStack[V]) Pop() (V, error) {
	if this.i == 0 {
		return *new(V), EmptyStackError
	}

	res := this.s[this.i-1]
	this.i--

	return res, nil
}

func (this SizedStack[V]) Len() int {
	return this.i
}

func (this SizedStack[V]) Cap() int {
	return this.cap
}
