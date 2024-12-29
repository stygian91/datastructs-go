package set_test

import (
	"slices"
	"testing"

	"github.com/stygian91/datastructs-go/set"
)

func TestSet(t *testing.T) {
	s := set.New[int]()

	exists := s.Add(42)
	if s.Len() != 1 {
		t.Fatalf("Expected len: %d, got: %d", 1, s.Len())
	}
	if exists {
		t.Fatal("Expected 42 to not exist on first Add(42)")
	}

	if s.Len() != 1 {
		t.Fatalf("Expected len: %d, got: %d", 1, s.Len())
	}
	exists = s.Add(42)
	if !exists {
		t.Fatal("Expected 42 to exist on second Add(42)")
	}

	exists = s.Add(69)
	if s.Len() != 2 {
		t.Fatalf("Expected len: %d, got: %d", 2, s.Len())
	}
	if exists {
		t.Fatal("Expected 69 to not exist on first Add(69)")
	}

	existed := s.Remove(42)
	if !existed {
		t.Fatal("Expected 42 to exist on first Remove(42)")
	}
	if s.Len() != 1 {
		t.Fatalf("Expected len: %d, got: %d", 1, s.Len())
	}

	existed = s.Remove(42)
	if existed {
		t.Fatal("Expected 42 to not exist on second Remove(42)")
	}
	if s.Len() != 1 {
		t.Fatalf("Expected len: %d, got: %d", 1, s.Len())
	}

	existed = s.Remove(69)
	if !existed {
		t.Fatal("Expected 69 to exist on first Remove(69)")
	}
	if s.Len() != 0 {
		t.Fatalf("Expected len: %d, got: %d", 0, s.Len())
	}
}

func TestSeq(t *testing.T) {
	s := set.New[int]()
	values := []int{1, 2, 3, 42, 69}

	for _, v := range values {
		s.Add(v)
	}

	cnt := 0
	for v := range s.Seq() {
		if !slices.Contains(values, v) {
			t.Fatalf("Exptected to find %d", v)
		}

		cnt++
	}

	if cnt != len(values) {
		t.Fatalf("Expected Seq() len to be: %d, got: %d", len(values), cnt)
	}
}

func TestMerge(t *testing.T) {
	sa := set.New[int]()
	sb := set.New[int]()
	a := []int{1, 2, 3, 6}
	b := []int{3, 4, 5, 6}

	for _, v := range a {
		sa.Add(v)
	}

	for _, v := range b {
		sb.Add(v)
	}

	sa.Merge(sb)

	if sa.Len() != 6 {
		t.Errorf("Merge len mismatch - expected: %d, got %d\n", 6, sa.Len())
	}

	exp := []int{1, 2, 3, 4, 5, 6}
	for _, v := range exp {
		if !sa.Contains(v) {
			t.Errorf("Value %d not found in set after merge\n", v)
		}
	}
}
