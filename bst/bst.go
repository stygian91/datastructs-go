package bst

import (
	"iter"

	"golang.org/x/exp/constraints"
)

type ord constraints.Ordered

type BSTNode[T ord] struct {
	Value T
	Left  *BSTNode[T]
	Right *BSTNode[T]
}

type BST[T ord] struct {
	Root BSTNode[T]
}

func NewBST[T ord](value T) BST[T] {
	return BST[T]{Root: NewBSTNode(value)}
}

func NewBSTNode[T ord](value T) BSTNode[T] {
	return BSTNode[T]{Value: value}
}

func (this *BST[T]) Add(value T)                         { this.Root.Add(value) }
func (this *BST[T]) InOrderSeq() iter.Seq[*BSTNode[T]]   { return this.Root.InOrderSeq() }
func (this *BST[T]) PostOrderSeq() iter.Seq[*BSTNode[T]] { return this.Root.PostOrderSeq() }
func (this *BST[T]) Search(value T) (*BSTNode[T], bool)  { return this.Root.Search(value) }
func (this *BST[T]) Min() *BSTNode[T]                    { return this.Root.Min() }
func (this *BST[T]) Max() *BSTNode[T]                    { return this.Root.Max() }

func (this *BSTNode[T]) Add(value T) {
	curr := this

	for curr.Value != value {
		if value < curr.Value {
			if curr.Left != nil {
				curr = curr.Left
				continue
			}

			newLeft := NewBSTNode(value)
			curr.Left = &newLeft
			break
		}

		if curr.Right != nil {
			curr = curr.Right
			continue
		}

		newRight := NewBSTNode(value)
		curr.Right = &newRight
		break
	}
}

func (this *BSTNode[T]) Min() *BSTNode[T] {
	curr := this

	for {
		if curr.Left == nil {
			return curr
		}

		curr = curr.Left
	}
}

func (this *BSTNode[T]) Max() *BSTNode[T] {
	curr := this

	for {
		if curr.Right == nil {
			return curr
		}

		curr = curr.Right
	}
}

func (this *BSTNode[T]) InOrderSeq() iter.Seq[*BSTNode[T]] {
	return func(yield func(*BSTNode[T]) bool) {
		curr := this
		s := []*BSTNode[T]{}

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

func (this *BSTNode[T]) PostOrderSeq() iter.Seq[*BSTNode[T]] {
	return func(yield func(*BSTNode[T]) bool) {
		curr := this
		s := []*BSTNode[T]{}

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

func (this *BSTNode[T]) Search(value T) (*BSTNode[T], bool) {
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
