package slices_test

import (
	"testing"

	"github.com/stygian91/datastructs-go/slices"
)

func TestLastIndex(t *testing.T) {
	data := []int{1, 3, 4, 1, 2}
	test := func(data []int, search, expected int) {
		res := slices.LastIndex(data, search)
		if res != expected {
			t.Errorf("Expected: %d, got %d\n", expected, res)
		}
	}

	test(data, 2, 4)
	test(data, 1, 3)
	test(data, 0, -1)
	test([]int{}, 1, -1)
}

func TestLastIndexFunc(t *testing.T) {
	data := []int{1, 2, 3, 5, 1, 4}
	res1 := slices.LastIndexFunc(data, func(a int) bool { return a < 4 })
	exp1 := 4
	if res1 != exp1 {
		t.Errorf("Expected: %d, got %d\n", exp1, res1)
	}

	res2 := slices.LastIndexFunc(data, func(a int) bool { return a > 5 })
	exp2 := -1
	if res2 != exp2 {
		t.Errorf("Expected: %d, got %d\n", exp2, res2)
	}

	res3 := slices.LastIndexFunc([]int{}, func(a int) bool { return a < 4 })
	exp3 := -1
	if res3 != exp3 {
		t.Errorf("Expected: %d, got %d\n", exp3, res3)
	}
}
