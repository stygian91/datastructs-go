package bst_test

import (
	"testing"

	"github.com/stygian91/datastructs-go/bst"
)

func TestBST(t *testing.T) {
	tree := bst.NewBST(3)
	tree.Add(2)
	tree.Add(4)
	tree.Add(1)
	tree.Add(5)

	i := 1
	for node := range tree.InOrderSeq() {
		if node.Value != i {
			t.Errorf("InOrderSeq(): expected %d, got %d", i, node.Value)
			return
		}
		i++
	}

	i = 5
	for node := range tree.PostOrderSeq() {
		if node.Value != i {
			t.Errorf("PostOrderSeq(): expected %d, got %d", i, node.Value)
			return
		}
		i--
	}

	expected := []int{3, 2, 1, 4, 5}
	i = 0
	for node := range tree.PreOrderSeq() {
		if node.Value != expected[i] {
			t.Errorf("PreOrderSeq(): expected %d, got %d", expected[i], node.Value)
			return
		}
		i++
	}

	min := tree.Min().Value
	if min != 1 {
		t.Errorf("Min(): expected 1, got %d", min)
		return
	}

	max := tree.Max().Value
	if max != 5 {
		t.Errorf("Max(): expected 5, got %d", max)
		return
	}

	if found, exists := tree.Search(2); !exists || found.Value != 2 {
		t.Error("Search(): expected to find 2")
		return
	}

	if _, exists := tree.Search(6); exists {
		t.Error("Search(): expected to not find 6")
		return
	}
}

func TestBSTBalanced(t *testing.T) {
	tree := bst.NewBST(0)

	for i := 1; i < 10; i++ {
		tree.Add(i)
	}

	balanced := tree.NewBalanced()
	expected := []int{4, 1, 0, 2, 3, 7, 5, 6, 8, 9}

	i := 0
	for node := range balanced.PreOrderSeq() {
		if node.Value != expected[i] {
			t.Errorf("TestBSTBalanced(): expected %d, got %d", expected[i], node.Value)
			return
		}
		i++
	}
}
