package bst

import (
	"iter"

	st "github.com/stygian91/datastructs-go/stack"
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

	root := FromSortedList(values)

	return BST[T]{Root: root}
}

func FromSortedList[T ordered](values []T) Node[T] {
	if len(values) == 0 {
		return Node[T]{}
	}

	type Frame struct {
		start, end int
		node       *Node[T]
		isLeft     bool
	}

	start := 0
	end := len(values) - 1
	mid := (start + end) / 2
	root := NewNode(values[mid])
	s := st.NewStack[Frame]()
	s.Push(
		Frame{start: mid + 1, end: end, node: &root, isLeft: false},
		Frame{start: start, end: mid - 1, node: &root, isLeft: true},
	)

	for s.Len() > 0 {
		f, _ := s.Pop()

		if f.start > f.end {
			continue
		}

		mid = (f.start + f.end) / 2
		node := NewNode(values[mid])
		if f.isLeft {
			f.node.Left = &node
		} else {
			f.node.Right = &node
		}

		s.Push(
			Frame{start: mid + 1, end: f.end, node: &node, isLeft: false},
			Frame{start: f.start, end: mid - 1, node: &node, isLeft: true},
		)
	}

	return root
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
		s := st.NewStack[*Node[T]]()

		for curr != nil || s.Len() > 0 {
			for curr != nil {
				s.Push(curr)
				curr = curr.Left
			}

			curr, _ = s.Pop()

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
		s := st.NewStack[*Node[T]]()

		for curr != nil || s.Len() > 0 {
			for curr != nil {
				s.Push(curr)
				curr = curr.Right
			}

			curr, _ = s.Pop()

			if !yield(curr) {
				return
			}

			curr = curr.Left
		}
	}
}

func (this *Node[T]) PreOrderSeq() iter.Seq[*Node[T]] {
	return func(yield func(*Node[T]) bool) {
		s := st.NewStack[*Node[T]]()

		for s.Len() > 0 {
			curr, _ := s.Pop()

			if curr == nil {
				continue
			}

			if !yield(curr) {
				return
			}

			s.Push(curr.Right, curr.Left)
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
