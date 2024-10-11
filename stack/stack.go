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

type LimitedStack[V any] struct {
	s   []V
	cap int
}

var StackOverflowError = fmt.Errorf("Trying to push to a full stack.")

func NewLimitedStack[V any](cap int) LimitedStack[V] {
	return LimitedStack[V]{cap: cap, s: []V{}}
}

func (this *LimitedStack[V]) Push(vals ...V) error {
	if len(this.s) + len(vals) > this.cap {
		return StackOverflowError
	}

	this.s = append(this.s, vals...)

	return nil
}

func (this *LimitedStack[V]) Pop() (V, error) {
	if len(this.s) == 0 {
		return *new(V), EmptyStackError
	}

	res := this.s[len(this.s)-1]
	this.s = this.s[:len(this.s)-1]

	return res, nil
}

func (this LimitedStack[V]) Len() int {
	return len(this.s)
}

func (this LimitedStack[V]) Cap() int {
	return this.cap
}
