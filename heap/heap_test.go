package heap_test

import (
	"testing"

	"github.com/stygian91/datastructs-go/heap"
)

func TestHeap(t *testing.T) {
	h := heap.New[int]()
	h.Push(5)
	h.Push(3)
	h.Push(6)
	h.Push(2)
	h.Push(1)

	if h.Len() != 5 {
		t.Fatalf("Expected len: %d, got %d\n", 5, h.Len())
	}

	exp := []int{1, 2, 3, 5, 6}
	for _, e := range exp {
		r := h.Pop()
		if r != e {
			t.Errorf("Expected %d, got %d\n", e, r)
		}
	}
}
