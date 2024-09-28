package bst

import (
	"iter"

	"golang.org/x/exp/constraints"
)

type ordered constraints.Ordered

type Node[T ordered] struct {
	Value T
	Left  *Node[T]
	Right *Node[T]
}

type BST[T ordered] struct {
	Root Node[T]
}

func NewBST[T ordered](value T) BST[T] {
	return BST[T]{Root: NewNode(value)}
}

func NewNode[T ordered](value T) Node[T] {
	return Node[T]{Value: value}
}

func (this *BST[T]) Add(value T)                      { this.Root.Add(value) }
func (this *BST[T]) InOrderSeq() iter.Seq[*Node[T]]   { return this.Root.InOrderSeq() }
func (this *BST[T]) PostOrderSeq() iter.Seq[*Node[T]] { return this.Root.PostOrderSeq() }
func (this *BST[T]) PreOrderSeq() iter.Seq[*Node[T]]  { return this.Root.PreOrderSeq() }
func (this *BST[T]) Search(value T) (*Node[T], bool)  { return this.Root.Search(value) }
func (this *BST[T]) Min() *Node[T]                    { return this.Root.Min() }
func (this *BST[T]) Max() *Node[T]                    { return this.Root.Max() }

func (this BST[T]) NewBalanced() BST[T] {
	values := []T{}
	for node := range this.InOrderSeq() {
		values = append(values, node.Value)
	}

	root, _ := FromSortedList(values, 0, len(values)-1)

	return BST[T]{Root: root}
}

func FromSortedList[T ordered](values []T, start, end int) (Node[T], bool) {
	if start > end {
		return Node[T]{}, false
	}

	mid := (start + end) / 2
	root := NewNode(values[mid])

	if left, lexists := FromSortedList(values, start, mid-1); lexists {
		root.Left = &left
	}

	if right, rexists := FromSortedList(values, mid+1, end); rexists {
		root.Right = &right
	}

	return root, true
}

func (this *Node[T]) Add(value T) {
	curr := this

	for curr.Value != value {
		if value < curr.Value {
			if curr.Left != nil {
				curr = curr.Left
				continue
			}

			newLeft := NewNode(value)
			curr.Left = &newLeft
			break
		}

		if curr.Right != nil {
			curr = curr.Right
			continue
		}

		newRight := NewNode(value)
		curr.Right = &newRight
		break
	}
}

func (this *Node[T]) Min() *Node[T] {
	curr := this

	for {
		if curr.Left == nil {
			return curr
		}

		curr = curr.Left
	}
}

func (this *Node[T]) Max() *Node[T] {
	curr := this

	for {
		if curr.Right == nil {
			return curr
		}

		curr = curr.Right
	}
}

func (this *Node[T]) InOrderSeq() iter.Seq[*Node[T]] {
	return func(yield func(*Node[T]) bool) {
		curr := this
		s := []*Node[T]{}

		for curr != nil || len(s) > 0 {
			for curr != nil {
				s = append(s, curr)
				curr = curr.Left
			}

			curr = s[len(s)-1]
			s = s[:len(s)-1]

			if !yield(curr) {
				return
			}

			curr = curr.Right
		}
	}
}

func (this *Node[T]) PostOrderSeq() iter.Seq[*Node[T]] {
	return func(yield func(*Node[T]) bool) {
		curr := this
		s := []*Node[T]{}

		for curr != nil || len(s) > 0 {
			for curr != nil {
				s = append(s, curr)
				curr = curr.Right
			}

			curr = s[len(s)-1]
			s = s[:len(s)-1]

			if !yield(curr) {
				return
			}

			curr = curr.Left
		}
	}
}

func (this *Node[T]) PreOrderSeq() iter.Seq[*Node[T]] {
	return func(yield func(*Node[T]) bool) {
		s := []*Node[T]{this}

		for len(s) > 0 {
			curr := s[len(s)-1]
			s = s[:len(s)-1]

			if curr == nil {
				continue
			}

			if !yield(curr) {
				return
			}

			s = append(s, curr.Right, curr.Left)
		}
	}
}

func (this *Node[T]) Search(value T) (*Node[T], bool) {
	curr := this

	for {
		if curr.Value == value {
			return curr, true
		}

		if value < curr.Value {
			if curr.Left != nil {
				curr = curr.Left
				continue
			}

			return nil, false
		}

		if curr.Right != nil {
			curr = curr.Right
			continue
		}

		return nil, false
	}
}
