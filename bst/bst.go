package bst

import (
	"iter"

	st "github.com/stygian91/datastructs-go/stack"
	"golang.org/x/exp/constraints"
)

type ordered constraints.Ordered

type NodeValue[T ordered, M any] struct {
	Value T
	Meta  M
}

type Node[T ordered, M any] struct {
	Value T
	Meta  M
	Left  *Node[T, M]
	Right *Node[T, M]
}

type BST[T ordered, M any] struct {
	Root Node[T, M]
}

func NewBST[T ordered, M any](value T, meta M) BST[T, M] {
	return BST[T, M]{Root: NewNode(value, meta)}
}

func NewNode[T ordered, M any](value T, meta M) Node[T, M] {
	return Node[T, M]{Value: value, Meta: meta}
}

func (this *BST[T, M]) Add(value T, meta M)                 { this.Root.Add(value, meta) }
func (this *BST[T, M]) InOrderSeq() iter.Seq[*Node[T, M]]   { return this.Root.InOrderSeq() }
func (this *BST[T, M]) PostOrderSeq() iter.Seq[*Node[T, M]] { return this.Root.PostOrderSeq() }
func (this *BST[T, M]) PreOrderSeq() iter.Seq[*Node[T, M]]  { return this.Root.PreOrderSeq() }
func (this *BST[T, M]) Search(value T) (*Node[T, M], bool)  { return this.Root.Search(value) }
func (this *BST[T, M]) Min() *Node[T, M]                    { return this.Root.Min() }
func (this *BST[T, M]) Max() *Node[T, M]                    { return this.Root.Max() }

func (this BST[T, M]) NewBalanced() BST[T, M] {
	values := []NodeValue[T, M]{}
	for node := range this.InOrderSeq() {
		values = append(values, NodeValue[T, M]{Value: node.Value, Meta: node.Meta})
	}

	root := FromSortedList(values)

	return BST[T, M]{Root: root}
}

func FromSortedList[T ordered, M any](values []NodeValue[T, M]) Node[T, M] {
	if len(values) == 0 {
		return Node[T, M]{}
	}

	type Frame struct {
		start, end int
		node       *Node[T, M]
		isLeft     bool
	}

	start := 0
	end := len(values) - 1
	mid := (start + end) / 2
	root := NewNode(values[mid].Value, values[mid].Meta)
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
		node := NewNode(values[mid].Value, values[mid].Meta)
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

func (this *Node[T, M]) Add(value T, meta M) {
	curr := this

	for curr.Value != value {
		if value < curr.Value {
			if curr.Left != nil {
				curr = curr.Left
				continue
			}

			newLeft := NewNode(value, meta)
			curr.Left = &newLeft
			break
		}

		if curr.Right != nil {
			curr = curr.Right
			continue
		}

		newRight := NewNode(value, meta)
		curr.Right = &newRight
		break
	}
}

func (this *BST[T, M]) Remove(value T) (*Node[T, M], bool) {
	values := []NodeValue[T, M]{}
	found := false
	foundNode := &Node[T, M]{}

	for node := range this.InOrderSeq() {
		if !found && node.Value == value {
			found = true
			foundNode = node
			continue
		}

		values = append(values, NodeValue[T, M]{Value: node.Value, Meta: node.Meta})
	}

	if found {
		root := FromSortedList(values)
		this.Root = root
	}

	return foundNode, found
}

func (this *Node[T, M]) Min() *Node[T, M] {
	curr := this

	for {
		if curr.Left == nil {
			return curr
		}

		curr = curr.Left
	}
}

func (this *Node[T, M]) Max() *Node[T, M] {
	curr := this

	for {
		if curr.Right == nil {
			return curr
		}

		curr = curr.Right
	}
}

func (this *Node[T, M]) InOrderSeq() iter.Seq[*Node[T, M]] {
	return func(yield func(*Node[T, M]) bool) {
		curr := this
		s := st.NewStack[*Node[T, M]]()

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

func (this *Node[T, M]) PostOrderSeq() iter.Seq[*Node[T, M]] {
	return func(yield func(*Node[T, M]) bool) {
		curr := this
		s := st.NewStack[*Node[T, M]]()

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

func (this *Node[T, M]) PreOrderSeq() iter.Seq[*Node[T, M]] {
	return func(yield func(*Node[T, M]) bool) {
		s := st.NewStack[*Node[T, M]]()

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

func (this *Node[T, M]) Search(value T) (*Node[T, M], bool) {
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
